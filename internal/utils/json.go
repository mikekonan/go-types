package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"unsafe"
)

const (
	nullBytes        = "null"
	emptyStringBytes = `""`
)

var (
	reflectTypeOfString = reflect.TypeOf("")
)

// UnsafeStringFromJson converts a JSON string literal in data to a Go string.
// It treats the JSON values `null` and `""` as empty values and returns isEmptyValue = true.
// If data is not a JSON string literal (not enclosed in double quotes), it returns a *json.UnmarshalTypeError.
// On success it returns the unquoted string, isEmptyValue = false, and nil error.
func UnsafeStringFromJson(data []byte) (v string, isEmptyValue bool, err error) {
	// check for null or empty string
	if bytes.Equal(data, []byte(nullBytes)) || bytes.Equal(data, []byte(emptyStringBytes)) {
		return "", true, nil
	}

	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return "", false, &json.UnmarshalTypeError{Value: string(data), Type: reflectTypeOfString}
	}

	data = data[1 : len(data)-1]

	return unsafe.String(unsafe.SliceData(data), len(data)), false, nil
}
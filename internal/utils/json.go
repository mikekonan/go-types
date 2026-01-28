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

package language

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// Alpha3Code represents alpha-3 code
type Alpha3Code string

// UnmarshalJSON unmarshall implementation for alpha3code
func (code *Alpha3Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Alpha3Code(str)
	if len(enumValue) != 0 {
		if _, err := ByAlpha3CodeErr(enumValue); err != nil {
			return err
		}
	}

	*code = enumValue
	return nil
}

// Value implementation of driver.Valuer
func (code Alpha3Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	var language Language

	if language, err = ByAlpha3CodeErr(code); err != nil {
		return nil, err
	}

	return language.Alpha3Code().String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code Alpha3Code) Validate() (err error) {
	_, err = ByAlpha3CodeErr(code)

	return
}

// IsSet indicates if Name is set
func (code Alpha3Code) IsSet() bool {
	return len(string(code)) > 0
}

// String implementation of Stringer interface
func (code Alpha3Code) String() string {
	return strings.ToLower(string(code))
}

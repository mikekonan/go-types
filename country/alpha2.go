package country

import (
	"database/sql/driver"
	"encoding/json"
)

// Alpha2Code represents alpha-2 code
type Alpha2Code string

// UnmarshalJSON unmarshall implementation for alpha2code
func (code *Alpha2Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Alpha2Code(str)
	if _, err := ByAlpha2CodeErr(enumValue); err != nil {
		return err
	}

	*code = enumValue
	return nil
}

// Value implementation of driver.Valuer
func (code Alpha2Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	var country Country

	if country, err = ByAlpha2CodeErr(code); err != nil {
		return nil, err
	}

	return country.Alpha2Code().String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code Alpha2Code) Validate() (err error) {
	_, err = ByAlpha2CodeErr(code)

	return
}

// IsSet indicates if Name is set
func (code Alpha2Code) IsSet() bool {
	return len(string(code)) > 0
}

// String implementation of Stringer interface
func (code Alpha2Code) String() string {
	return string(code)
}

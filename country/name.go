package country

import (
	"database/sql/driver"
	"encoding/json"
)

// Name represents country name
type Name string

// UnmarshalJSON unmarshall implementation for name
func (name *Name) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Name(str)
	if len(enumValue) != 0 {
		if _, err := ByNameErr(enumValue); err != nil {
			return err
		}
	}

	*name = enumValue
	return nil
}

// Value implementation of driver.Valuer
func (name Name) Value() (value driver.Value, err error) {
	if name == "" {
		return "", nil
	}

	if _, err = ByNameErr(name); err != nil {
		return nil, err
	}

	return name.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (name Name) Validate() (err error) {
	_, err = ByNameErr(name)

	return
}

// IsSet indicates if Name is set
func (name Name) IsSet() bool {
	return len(string(name)) > 0
}

// String implementation of Stringer interface
func (name Name) String() string {
	return string(name)
}

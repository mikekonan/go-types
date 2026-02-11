package subdivision

import (
	"database/sql/driver"
	"encoding/json"
)

// Name represents subdivision name
type Name string

// UnmarshalJSON unmarshall implementation for name
func (name *Name) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*name = Name(str)
	return nil
}

// Value implementation of driver.Valuer
func (name Name) Value() (value driver.Value, err error) {
	return name.String(), nil
}

// IsSet indicates if Name is set
func (name Name) IsSet() bool {
	return len(string(name)) > 0
}

// String implementation of Stringer interface
func (name Name) String() string {
	return string(name)
}

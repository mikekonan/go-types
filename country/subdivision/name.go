package subdivision

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Name represents subdivision name
type Name string

// UnmarshalJSON unmarshall implementation for name
func (name *Name) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == "" {
		return nil
	}

	if _, ok := subdivisionNames[str]; !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-2 subdivision name", str)
	}

	*name = Name(str)
	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (name Name) Validate() error {
	if _, ok := subdivisionNames[string(name)]; !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-2 subdivision name", name)
	}

	return nil
}

// Value implementation of driver.Valuer
func (name Name) Value() (value driver.Value, err error) {
	if name == "" {
		return "", nil
	}

	if err = name.Validate(); err != nil {
		return nil, err
	}

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

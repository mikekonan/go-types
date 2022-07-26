package currency

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Code represents a code type from ISO-4217
type Code string

// Value implementation of driver.Valuer
func (code Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	if err = code.Validate(); err != nil {
		return nil, err
	}

	return code.String(), nil
}

// UnmarshalJSON unmarshall implementation for Code
func (code *Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	currency, err := ByCodeStrErr(str)
	if err != nil {
		return err
	}

	*code = currency.Code()

	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (code Code) Validate() error {
	if _, ok := ByCodeStr(string(code)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return nil
}

// IsSet indicates if Code is set
func (code Code) IsSet() bool {
	return len(string(code)) > 0
}

// String implementation of Stringer interface
func (code Code) String() string {
	return string(code)
}

// Places returns number of digits after the dot. E.g. 2 for USD, 0 for currencies which do not support fractions
func (code Code) Places() int {
	curr, _ := ByCodeStr(string(code))

	return curr.DecimalPlaces()
}

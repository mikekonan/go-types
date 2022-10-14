package phone

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Number represents a phone number type
type Number string

// Value implementation of driver.Valuer
func (number Number) Value() (value driver.Value, err error) {
	if number == "" {
		return "", nil
	}

	if err = number.Validate(); err != nil {
		return nil, err
	}

	return number.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (number Number) Validate() error {
	var err error
	if _, err = strconv.Atoi(number.String()); err != nil {
		return fmt.Errorf("'%s' is not valid phone number", number)
	}

	// minimum length is 7 for Saint Helena (Format: +290 XXXX) and Niue (Format: +683 XXXX)
	if len(number) < 7 || len(number) > 15 {
		return fmt.Errorf("'%s' is not valid phone number", number)
	}

	return nil
}

// UnmarshalJSON unmarshall implementation
func (number *Number) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	phoneValue := Number(str)
	if len(phoneValue) != 0 {
		if err := phoneValue.Validate(); err != nil {
			return err
		}
	}

	*number = phoneValue

	return nil
}

// String implementation of Stringer interface
func (number Number) String() string {
	return string(number)
}

package postalcode

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

var stringValid = regexp.MustCompile(`^[0-9A-Za-z\s\-]+$`).MatchString

// PostalCode represents a phone number type
type PostalCode string

// Value implementation of driver.Valuer
func (code PostalCode) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	if err = code.Validate(); err != nil {
		return nil, err
	}

	return code.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code PostalCode) Validate() error {
	if len(code) < 2 || len(code) > 10 {
		return fmt.Errorf("'%s' is not valid postal code", code)
	}

	if !stringValid(code.String()) {
		return errors.New("postal code should contain numbers, letters, spaces and hyphen only")
	}

	return nil
}

// UnmarshalJSON unmarshall implementation
func (code *PostalCode) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	codeValue := PostalCode(str)
	if len(codeValue) != 0 {
		if err := codeValue.Validate(); err != nil {
			return err
		}
	}

	*code = codeValue

	return nil
}

// String implementation of Stringer interface
func (code PostalCode) String() string {
	return string(code)
}

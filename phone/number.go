package phone

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/mikekonan/go-types/v2/country"
	"strconv"

	"github.com/nyaruka/phonenumbers"
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

// StrictParse is Parse analogue with phone number region consistency check: fails if passed country is not matches parsed one
func StrictParse(rawPhone string, country country.Alpha2Code) (DialCode, Number, error) {
	var parsed, err = phonenumbers.Parse(rawPhone, country.String())
	if err != nil {
		return "", "", err
	}

	var valid = phonenumbers.IsValidNumber(parsed)
	if !valid {
		return "", "", fmt.Errorf(`%s phone number is not valid for %s region`, rawPhone, country.String())
	}

	var dialCode = DialCode(strconv.Itoa(int(parsed.GetCountryCode())))
	if err = dialCode.Validate(); err != nil {
		return "", "", err
	}

	return dialCode, Number(strconv.FormatUint(parsed.GetNationalNumber(), 10)), nil
}

// Parse matches passed 'country' string representation to digits and searches for it at the beginning of 'rawPhone'.
// If not found - consider whole passes 'rawPhone' as Number and use matched one 'country' as DialCode
func Parse(rawPhone string, country country.Alpha2Code) (DialCode, Number, error) {
	var parsed, err = phonenumbers.Parse(rawPhone, country.String())
	if err != nil {
		return "", "", err
	}

	var dialCode = DialCode(strconv.Itoa(int(parsed.GetCountryCode())))
	if err = dialCode.Validate(); err != nil {
		return "", "", err
	}

	return dialCode, Number(strconv.FormatUint(parsed.GetNationalNumber(), 10)), nil
}

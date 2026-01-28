package currency

import (
	"database/sql/driver"

	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Number represents a number type from ISO-4217
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

// UnmarshalJSON unmarshall implementation for Number
func (number *Number) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		return newInvalidDataError(str, standardISO4217Number)
	}

	currency, err := ByNumberStrErr(str)
	if err != nil {
		return err
	}

	*number = currency.Number()

	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (number Number) Validate() error {
	if _, ok := ByNumberStr(string(number)); !ok {
		return newInvalidDataError(string(number), standardISO4217Number)
	}

	return nil
}

// IsSet indicates if Number is set
func (number Number) IsSet() bool {
	return len(string(number)) > 0
}

// String implementation of Stringer interface
func (number Number) String() string {
	return string(number)
}

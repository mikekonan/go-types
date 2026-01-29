package country

import (
	"database/sql/driver"

	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Alpha2Code represents alpha-2 code
type Alpha2Code string

// UnmarshalJSON unmarshall implementation for alpha2code
func (code *Alpha2Code) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		// empty value is allowed for Alpha2Code
		return nil
	}

	country, err := ByAlpha2CodeErr(Alpha2Code(str))
	if err != nil {
		return err
	}

	*code = country.Alpha2Code()

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

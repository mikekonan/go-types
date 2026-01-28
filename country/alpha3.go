package country

import (
	"database/sql/driver"

	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Alpha3Code represents alpha-3 code
type Alpha3Code string

// UnmarshalJSON unmarshall implementation for alpha3code
func (code *Alpha3Code) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		return nil
	}

	country, err := ByAlpha3CodeErr(Alpha3Code(str))
	if err != nil {
		return err
	}

	*code = country.Alpha3Code()

	return nil
}

// Value implementation of driver.Valuer
func (code Alpha3Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	var country Country

	if country, err = ByAlpha3CodeErr(code); err != nil {
		return nil, err
	}

	return country.Alpha3Code().String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code Alpha3Code) Validate() (err error) {
	_, err = ByAlpha3CodeErr(code)

	return
}

// IsSet indicates if Name is set
func (code Alpha3Code) IsSet() bool {
	return len(string(code)) > 0
}

// String implementation of Stringer interface
func (code Alpha3Code) String() string {
	return string(code)
}

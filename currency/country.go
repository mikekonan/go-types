package currency

import (
	"database/sql/driver"

	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Country represents a country type from ISO-4217
type Country string

// Value implementation of driver.Valuer
func (country Country) Value() (value driver.Value, err error) {
	if country == "" {
		return "", nil
	}

	if err = country.Validate(); err != nil {
		return nil, err
	}

	return country.String(), nil
}

// UnmarshalJSON unmarshall implementation for Country
func (country *Country) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		return newInvalidDataError(str, standardISO4217Country)
	}

	_, err = ByCountryStrErr(str)
	if err != nil {
		return err
	}

	*country = Country(str)

	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (country Country) Validate() error {
	if _, ok := ByCountryStr(string(country)); !ok {
		return newInvalidDataError(string(country), standardISO4217Country)
	}

	return nil
}

// IsSet indicates if Country is set
func (country Country) IsSet() bool {
	return len(string(country)) > 0
}

// String implementation of Stringer interface
func (country Country) String() string {
	return string(country)
}

// Countries is a Country slice type
type Countries []Country

// IsCountryIn Checks there is a country in Countries
func (countries Countries) IsCountryIn(country string) bool {
	for _, c := range countries {
		if string(c) == country {
			return true
		}
	}

	return false
}

package currency

import (
	"bytes"
	"database/sql/driver"
	"unsafe"
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
	data = bytes.TrimPrefix(bytes.TrimSuffix(data, []byte("\"")), []byte("\""))
	var str = unsafe.String(unsafe.SliceData(data), len(data))

	_, err := ByCountryStrErr(str)
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

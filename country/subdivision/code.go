package subdivision

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/mikekonan/go-types/v2/country"
	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Code represents ISO 3166-2 subdivision code (e.g. "US-CA")
type Code string

// UnmarshalJSON unmarshall implementation for subdivision code
func (code *Code) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		return nil
	}

	subdivision, err := ByCodeErr(Code(str))
	if err != nil {
		return err
	}

	*code = subdivision.Code()

	return nil
}

// Value implementation of driver.Valuer
func (code Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	var subdivision Subdivision

	if subdivision, err = ByCodeErr(code); err != nil {
		return nil, err
	}

	return subdivision.Code().String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code Code) Validate() (err error) {
	_, err = ByCodeErr(code)

	return
}

// ValidateForCountry validates that subdivision code is valid and belongs to the given country.
// It extracts the country prefix from the code, verifies it is a known country,
// compares it with the expected country, and checks that the subdivision exists.
func (code Code) ValidateForCountry(countryCode country.Alpha2Code) error {
	str := strings.ToUpper(code.String())

	// extract country prefix (2 chars before "-")
	idx := strings.IndexByte(str, '-')
	if idx != 2 {
		return fmt.Errorf("'%s' is not valid ISO-3166-2 subdivision code format", code)
	}

	prefix := country.Alpha2Code(str[:2])

	// verify prefix is a real country
	if _, err := country.ByAlpha2CodeErr(prefix); err != nil {
		return fmt.Errorf("'%s' contains unknown country prefix '%s'", code, prefix)
	}

	// compare countries
	if !strings.EqualFold(prefix.String(), countryCode.String()) {
		return fmt.Errorf("subdivision '%s' belongs to country '%s', not '%s'", code, prefix, countryCode)
	}

	// verify subdivision exists for this country
	if _, err := ByCodeErr(code); err != nil {
		return err
	}

	return nil
}

// ValidateForCountryStr validates that subdivision code string is valid and belongs to the given country
func (code Code) ValidateForCountryStr(countryCode string) error {
	return code.ValidateForCountry(country.Alpha2Code(countryCode))
}

// IsSet indicates if Code is set
func (code Code) IsSet() bool {
	return len(string(code)) > 0
}

// String implementation of Stringer interface
func (code Code) String() string {
	return string(code)
}

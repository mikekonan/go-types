package phone

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/mikekonan/go-types/v2/country"
	"strings"
)

// DialCode represents dial country code
type DialCode string

// UnmarshalJSON unmarshall implementation for dial country code
func (code *DialCode) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	dialCode := DialCode(str)
	if len(dialCode) != 0 {
		if _, err := CountriesByAlpha2CodeErr(dialCode); err != nil {
			return err
		}
	}

	*code = dialCode
	return nil
}

// Value implementation of driver.Valuer
func (code DialCode) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	if _, err = CountriesByAlpha2CodeErr(code); err != nil {
		return nil, err
	}

	return code.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (code DialCode) Validate() (err error) {
	_, err = CountriesByAlpha2CodeErr(code)

	return
}

// String implementation of Stringer interface
func (code DialCode) String() string {
	return string(code)
}

// DialByAlpha2Code lookup for dial country code by alpha-2 country code
func DialByAlpha2Code(countryCode country.Alpha2Code) (result DialCode, ok bool) {
	result, ok = dialCodeByCountryAlpha2Str[strings.ToUpper(countryCode.String())]
	return
}

// DialByAlpha2CodeStr lookup for dial country code by alpha-2 country code string
func DialByAlpha2CodeStr(code string) (DialCode, bool) {
	return DialByAlpha2Code(country.Alpha2Code(code))
}

// DialByAlpha2CodeErr lookup for dial country code by alpha-2 code with error return type
func DialByAlpha2CodeErr(countryCode country.Alpha2Code) (result DialCode, err error) {
	var ok bool
	result, ok = DialByAlpha2Code(countryCode)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO 639-1 code", countryCode)
	}

	return
}

// DialByAlpha2CodeStrErr lookup for dial country code by alpha-2 country code string with error return type
func DialByAlpha2CodeStrErr(code string) (DialCode, error) {
	return DialByAlpha2CodeErr(country.Alpha2Code(code))
}

// CountriesByDialCode lookup for country alpha-2 code by dial country code
func CountriesByDialCode(dialCode DialCode) (result []country.Alpha2Code, ok bool) {
	result, ok = countryCodeByDialString[dialCode.String()]
	return
}

// CountriesByDialCodeStr lookup for country alpha-2 code by dial country code string
func CountriesByDialCodeStr(dialCode string) ([]country.Alpha2Code, bool) {
	return CountriesByDialCode(DialCode(dialCode))
}

// CountriesByAlpha2CodeErr lookup for country alpha-2 code by dial country code with error return type
func CountriesByAlpha2CodeErr(dialCode DialCode) (result []country.Alpha2Code, err error) {
	var ok bool
	result, ok = CountriesByDialCode(dialCode)
	if !ok {
		err = fmt.Errorf("'%s' is not valid dial country code", dialCode)
	}

	return
}

// CountriesByDialCodeStrErr lookup for country alpha-2 code by dial country code string with error return type
func CountriesByDialCodeStrErr(dialCode string) ([]country.Alpha2Code, error) {
	return CountriesByAlpha2CodeErr(DialCode(dialCode))
}

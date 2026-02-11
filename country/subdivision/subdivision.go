/*
Package subdivision provides lookup methods over ISO 3166-2 subdivisions.

Constants have been generated with code generation ./generator.

The source of truth is https://www.iso.org/iso-3166-country-codes.html.

Lookup methods by code are case robust.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package subdivision

import (
	"strings"

	"github.com/mikekonan/go-types/v2/country"
)

// Subdivision represents subdivision entity according to ISO 3166-2.
type Subdivision struct {
	name        Name
	code        Code
	countryCode country.Alpha2Code
	category    Category
}

// Name returns subdivision name
func (s Subdivision) Name() Name { return s.name }

// Code returns subdivision code
func (s Subdivision) Code() Code { return s.code }

// CountryCode returns country alpha-2 code
func (s Subdivision) CountryCode() country.Alpha2Code { return s.countryCode }

// Category returns subdivision category
func (s Subdivision) Category() Category { return s.category }

// NameStr returns subdivision name string
func (s Subdivision) NameStr() string { return s.name.String() }

// CodeStr returns subdivision code string
func (s Subdivision) CodeStr() string { return s.code.String() }

// CountryCodeStr returns country alpha-2 code string
func (s Subdivision) CountryCodeStr() string { return s.countryCode.String() }

// CategoryStr returns subdivision category string
func (s Subdivision) CategoryStr() string { return s.category.String() }

// ByCode lookup for subdivision by code
func ByCode(code Code) (result Subdivision, ok bool) {
	result, ok = SubdivisionByCode[strings.ToUpper(code.String())]

	return
}

// ByCodeStr lookup for subdivision by code string
func ByCodeStr(code string) (result Subdivision, ok bool) {
	return ByCode(Code(code))
}

// ByCodeErr lookup for subdivision by code with error return type
func ByCodeErr(code Code) (result Subdivision, err error) {
	var ok bool
	result, ok = ByCode(code)
	if !ok {
		err = newInvalidDataError(string(code), standardISO31662Code)
	}

	return
}

// ByCodeStrErr lookup for subdivision by code string with error return type
func ByCodeStrErr(code string) (result Subdivision, err error) {
	return ByCodeErr(Code(code))
}

// ByCountryCode lookup for subdivisions by country alpha-2 code
func ByCountryCode(code country.Alpha2Code) (result []Subdivision, ok bool) {
	result, ok = SubdivisionsByCountry[strings.ToUpper(code.String())]

	return
}

// ByCountryCodeStr lookup for subdivisions by country alpha-2 code string
func ByCountryCodeStr(code string) (result []Subdivision, ok bool) {
	return ByCountryCode(country.Alpha2Code(code))
}

// ByCountryCodeErr lookup for subdivisions by country alpha-2 code with error return type
func ByCountryCodeErr(code country.Alpha2Code) (result []Subdivision, err error) {
	var ok bool
	result, ok = ByCountryCode(code)
	if !ok {
		err = newInvalidDataError(string(code), standardISO31661Alpha2)
	}

	return
}

// ByCountryCodeStrErr lookup for subdivisions by country alpha-2 code string with error return type
func ByCountryCodeStrErr(code string) (result []Subdivision, err error) {
	return ByCountryCodeErr(country.Alpha2Code(code))
}

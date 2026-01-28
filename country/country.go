/*
Package country provides lookups methods over ISO-3166 countries.

Constants have been generated with code generation ./generator.

The source of truth is https://www.iso.org/iso-3166-country-codes.html.

Lookups methods by alpha2 and alpha3 are case robust.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package country

import (
	"strings"
)

// Country represents country entity according to ISO-3166.
type Country struct {
	name   Name
	alpha2 Alpha2Code
	alpha3 Alpha3Code
}

// Name returns country name
func (country Country) Name() Name { return country.name }

// Alpha2Code returns alpha-2 code
func (country Country) Alpha2Code() Alpha2Code { return country.alpha2 }

// Alpha3Code returns alpha-2 code
func (country Country) Alpha3Code() Alpha3Code { return country.alpha3 }

// NameStr returns country name string
func (country Country) NameStr() string { return country.name.String() }

// Alpha2CodeStr returns country alpha-2 code string
func (country Country) Alpha2CodeStr() string { return country.alpha2.String() }

// Alpha3CodeStr returns country alpha-2 code string
func (country Country) Alpha3CodeStr() string { return country.alpha3.String() }

// ByAlpha3Code lookup for country by alpha-3 code
func ByAlpha3Code(code Alpha3Code) (result Country, ok bool) {
	result, ok = CountryByAlpha3[strings.ToUpper(code.String())]

	return
}

// ByAlpha3CodeStr lookup for country by alpha-3 code string
func ByAlpha3CodeStr(code string) (result Country, ok bool) {
	return ByAlpha3Code(Alpha3Code(code))
}

// ByAlpha3CodeErr looks up a Country by its ISO 3166-1 alpha-3 code and returns an error if no match is found.
// Lookup is case-insensitive; when not found this returns a descriptive invalid-data error for the provided alpha-3 code.
func ByAlpha3CodeErr(code Alpha3Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha3Code(code)
	if !ok {
		err = newInvalidDataError(string(code), standardISO3166alpha3)
	}

	return
}

// ByAlpha3CodeStrErr lookup for country by alpha-3 code string with error return type
func ByAlpha3CodeStrErr(code string) (result Country, err error) {
	return ByAlpha3CodeErr(Alpha3Code(code))
}

// ByAlpha2Code lookup for country by alpha-2 code
func ByAlpha2Code(code Alpha2Code) (result Country, ok bool) {
	result, ok = CountryByAlpha2[strings.ToUpper(code.String())]
	return
}

// ByAlpha2CodeStr lookup for country by alpha-2 code string
func ByAlpha2CodeStr(code string) (result Country, ok bool) {
	return ByAlpha2Code(Alpha2Code(code))
}

// ByAlpha2CodeErr looks up a Country by its ISO 3166-1 alpha-2 code and returns an error when no match is found.
// If the lookup fails, the returned error identifies the invalid alpha-2 code.
func ByAlpha2CodeErr(code Alpha2Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha2Code(code)
	if !ok {
		err = newInvalidDataError(string(code), standardISO3166alpha2)
	}

	return
}

// ByAlpha2CodeStrErr lookup for country by alpha-2 code string with error return type
func ByAlpha2CodeStrErr(code string) (result Country, err error) {
	return ByAlpha2CodeErr(Alpha2Code(code))
}

// ByName lookup for country by name
func ByName(country Name) (result Country, ok bool) {
	result, ok = CountryByName[country.String()]
	return
}

// ByNameStr lookup for country by name string
func ByNameStr(country string) (result Country, ok bool) {
	return ByName(Name(country))
}

// ByNameErr looks up a country by its Name and returns the matching Country or an error if no match is found.
// If lookup fails, the returned error identifies the invalid country name according to the ISO-3166 country standard.
func ByNameErr(country Name) (result Country, err error) {
	var ok bool
	result, ok = ByName(country)
	if !ok {
		err = newInvalidDataError(string(country), standardISO3166Country)
	}

	return
}

// ByNameStrErr lookup for country by name string with error return type
func ByNameStrErr(country string) (result Country, err error) {
	return ByNameErr(Name(country))
}
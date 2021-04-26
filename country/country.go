/*
Package country provides lookups methods over ISO-3166 countries.

Constants have been generated with code generation ./generator.

The source of truth is https://www.iso.org/iso-3166-country-codes.html.

Lookups methods by alpha2 and alpha3 are case robust.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package country

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

//Name represents country name
type Name string

//UnmarshalJSON unmarshall implementation for name
func (name *Name) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Name(str)
	if _, err := ByNameErr(enumValue); err != nil {
		return err
	}

	*name = enumValue
	return nil
}

//Value implementation of driver.Valuer
func (name Name) Value() (value driver.Value, err error) {
	if name == "" {
		return "", nil
	}

	if _, err = ByNameErr(name); err != nil {
		return nil, err
	}

	return name.String(), nil
}

//Validate implementation of ozzo-validation Validate interface
func (name Name) Validate() (err error) {
	_, err = ByNameErr(name)

	return
}

//IsSet indicates if Name is set
func (name Name) IsSet() bool {
	return len(string(name)) > 0
}

//String implementation of Stringer interface
func (name Name) String() string {
	return string(name)
}

//Alpha2Code represents alpha-2 code
type Alpha2Code string

//UnmarshalJSON unmarshall implementation for alpha2code
func (code *Alpha2Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Alpha2Code(str)
	if _, err := ByAlpha2CodeErr(enumValue); err != nil {
		return err
	}

	*code = enumValue
	return nil
}

//Value implementation of driver.Valuer
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

//Validate implementation of ozzo-validation Validate interface
func (code Alpha2Code) Validate() (err error) {
	_, err = ByAlpha2CodeErr(code)

	return
}

//IsSet indicates if Name is set
func (code Alpha2Code) IsSet() bool {
	return len(string(code)) > 0
}

//String implementation of Stringer interface
func (code Alpha2Code) String() string {
	return string(code)
}

func (code Alpha2Code) toUpper() Alpha2Code {
	return Alpha2Code(strings.ToUpper(code.String()))
}

//Alpha3Code represents alpha-3 code
type Alpha3Code string

//UnmarshalJSON unmarshall implementation for alpha3code
func (code *Alpha3Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	enumValue := Alpha3Code(str)
	if _, err := ByAlpha3CodeErr(enumValue); err != nil {
		return err
	}

	*code = enumValue
	return nil
}

//Value implementation of driver.Valuer
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

//Validate implementation of ozzo-validation Validate interface
func (code Alpha3Code) Validate() (err error) {
	_, err = ByAlpha3CodeErr(code)

	return
}

//IsSet indicates if Name is set
func (code Alpha3Code) IsSet() bool {
	return len(string(code)) > 0
}

//String implementation of Stringer interface
func (code Alpha3Code) String() string {
	return string(code)
}

func (code Alpha3Code) toUpper() Alpha3Code {
	return Alpha3Code(strings.ToUpper(code.String()))
}

//Country represents country entity according to ISO-3166.
type Country struct {
	name   Name
	alpha2 Alpha2Code
	alpha3 Alpha3Code
}

//Name returns country name
func (country Country) Name() Name { return country.name }

//Alpha2Code returns alpha-2 code
func (country Country) Alpha2Code() Alpha2Code { return country.alpha2 }

//Alpha3Code returns alpha-2 code
func (country Country) Alpha3Code() Alpha3Code { return country.alpha3 }

//NameStr returns country name string
func (country Country) NameStr() string { return country.name.String() }

//Alpha2CodeStr returns country alpha-2 code string
func (country Country) Alpha2CodeStr() string { return country.alpha2.String() }

//Alpha3CodeStr returns country alpha-2 code string
func (country Country) Alpha3CodeStr() string { return country.alpha3.String() }

//ByAlpha3Code lookup for country by alpha-3 code
func ByAlpha3Code(code Alpha3Code) (result Country, ok bool) {
	result, ok = countryByAlpha3[code.toUpper()]

	return
}

//ByAlpha3CodeStr lookup for country by alpha-3 code string
func ByAlpha3CodeStr(code string) (result Country, ok bool) {
	return ByAlpha3Code(Alpha3Code(code))
}

//ByAlpha3CodeErr lookup for country by alpha-3 code with error return type
func ByAlpha3CodeErr(code Alpha3Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha3Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha3 code", code)
	}

	return
}

//ByAlpha3CodeStrErr lookup for country by alpha-3 code string with error return type
func ByAlpha3CodeStrErr(code string) (result Country, err error) {
	return ByAlpha3CodeErr(Alpha3Code(code))
}

//ByAlpha2Code lookup for country by alpha-2 code
func ByAlpha2Code(code Alpha2Code) (result Country, ok bool) {
	result, ok = countryByAlpha2[code.toUpper()]
	return
}

//ByAlpha2CodeStr lookup for country by alpha-2 code string
func ByAlpha2CodeStr(code string) (result Country, ok bool) {
	return ByAlpha2Code(Alpha2Code(code))
}

//ByAlpha2CodeErr lookup for country by alpha-2 code with error return type
func ByAlpha2CodeErr(code Alpha2Code) (result Country, err error) {
	var ok bool
	result, ok = ByAlpha2Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166-alpha2 code", code)
	}

	return
}

//ByAlpha2CodeStrErr lookup for country by alpha-2 code string with error return type
func ByAlpha2CodeStrErr(code string) (result Country, err error) {
	return ByAlpha2CodeErr(Alpha2Code(code))
}

//ByName lookup for country by name
func ByName(country Name) (result Country, ok bool) {
	result, ok = countryByName[country]
	return
}

//ByNameStr lookup for country by name string
func ByNameStr(country string) (result Country, ok bool) {
	return ByName(Name(country))
}

//ByNameErr lookup for country by name with error return type
func ByNameErr(country Name) (result Country, err error) {
	var ok bool
	result, ok = ByName(country)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO-3166 Country name", country)
	}

	return
}

//ByNameStrErr lookup for country by name string with error return type
func ByNameStrErr(country string) (result Country, err error) {
	return ByNameErr(Name(country))
}

/*
Package currency provides lookups methods over ISO-4217 currencies.

Constants have been generated with code generation ./generator.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package currency

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

//Country represents a country type from ISO-4217
type Country string

//Value implementation of driver.Valuer
func (country Country) Value() (value driver.Value, err error) {
	if country == "" {
		return "", nil
	}

	if err = country.Validate(); err != nil {
		return nil, err
	}

	return country.String(), nil
}

//UnmarshalJSON unmarshall implementation for Country
func (country *Country) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	_, err := ByCountryStrErr(str)
	if err != nil {
		return err
	}

	*country = Country(str)

	return nil
}

//Validate implementation of ozzo-validation Validate interface
func (country Country) Validate() error {
	if _, ok := ByCountryStr(string(country)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 country", country)
	}

	return nil
}

//IsSet indicates if Country is set
func (country Country) IsSet() bool {
	return len(string(country)) > 0
}

//String implementation of Stringer interface
func (country Country) String() string {
	return string(country)
}

//Countries is a Country slice type
type Countries []Country

//IsCountryIn Checks there is a country in Countries
func (countries Countries) IsCountryIn(country string) bool {
	for _, c := range countries {
		if string(c) == country {
			return true
		}
	}

	return false
}

//Currency represents a currency type from ISO-4217
type Currency string

//Value implementation of driver.Valuer
func (currency Currency) Value() (value driver.Value, err error) {
	if currency == "" {
		return "", nil
	}

	if err = currency.Validate(); err != nil {
		return nil, err
	}

	return currency.String(), nil
}

//UnmarshalJSON unmarshall implementation for Currency
func (currency *Currency) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	currencyValue, err := ByCurrencyStrErr(str)
	if err != nil {
		return err
	}

	*currency = currencyValue.Currency()

	return nil
}

//Validate implementation of ozzo-validation Validate interface
func (currency Currency) Validate() error {
	if _, ok := ByCurrencyStr(string(currency)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 currency", currency)
	}

	return nil
}

//IsSet indicates if Currency is set
func (currency Currency) IsSet() bool {
	return len(string(currency)) > 0
}

//String implementation of Stringer interface
func (currency Currency) String() string {
	return string(currency)
}

//Code represents a code type from ISO-4217
type Code string

//Value implementation of driver.Valuer
func (code Code) Value() (value driver.Value, err error) {
	if code == "" {
		return "", nil
	}

	if err = code.Validate(); err != nil {
		return nil, err
	}

	return code.String(), nil
}

//UnmarshalJSON unmarshall implementation for Code
func (code *Code) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	currency, err := ByCodeStrErr(str)
	if err != nil {
		return err
	}

	*code = currency.Code()

	return nil
}

//Validate implementation of ozzo-validation Validate interface
func (code Code) Validate() error {
	if _, ok := ByCodeStr(string(code)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return nil
}

//IsSet indicates if Code is set
func (code Code) IsSet() bool {
	return len(string(code)) > 0
}

//String implementation of Stringer interface
func (code Code) String() string {
	return string(code)
}

// Places returns number of digits after the dot. E.g. 2 for USD, 0 for for currencies which do not support fractions
func (code Code) Places() int {
	return currenciesByCode[string(code)].decimalPlaces
}

//Number represents a number type from ISO-4217
type Number string

//Value implementation of driver.Valuer
func (number Number) Value() (value driver.Value, err error) {
	if number == "" {
		return "", nil
	}

	if err = number.Validate(); err != nil {
		return nil, err
	}

	return number.String(), nil
}

//UnmarshalJSON unmarshall implementation for Number
func (number *Number) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	currency, err := ByNumberStrErr(str)
	if err != nil {
		return err
	}

	*number = currency.Number()

	return nil
}

//Validate implementation of ozzo-validation Validate interface
func (number Number) Validate() error {
	if _, ok := ByNumberStr(string(number)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 number", number)
	}

	return nil
}

//IsSet indicates if Number is set
func (number Number) IsSet() bool {
	return len(string(number)) > 0
}

//String implementation of Stringer interface
func (number Number) String() string {
	return string(number)
}

type currency struct {
	countries     Countries
	currency      Currency
	code          Code
	number        Number
	decimalPlaces int
}

type currencies []currency

//CurrencyByCurrency get currency by currency
func (currencies currencies) CurrencyByCurrency(curr string) (currency, bool) {
	for _, c := range currencies {
		if string(c.currency) == curr {
			return c, true
		}
	}

	return currency{}, false
}

//CurrencyByCode gets currency by code
func (currencies currencies) CurrencyByCode(code string) (currency, bool) {
	for _, c := range currencies {
		if string(c.code) == code {
			return c, true
		}
	}

	return currency{}, false
}

//CurrencyByNumber gets currency by number
func (currencies currencies) CurrencyByNumber(number string) (currency, bool) {
	for _, c := range currencies {
		if string(c.number) == number {
			return c, true
		}
	}

	return currency{}, false
}

//Currency returns Currency
func (c currency) Currency() Currency { return c.currency }

//Code returns Code
func (c currency) Code() Code { return c.code }

//Number returns Number
func (c currency) Number() Number { return c.number }

//Countries returns Countries
func (c currency) Countries() Countries { return c.countries }

//ByCodeStr lookup for currency type by code
func ByCodeStr(code string) (c currency, ok bool) {
	c, ok = currenciesByCode[code]
	return
}

//ByCurrencyStr lookup for currency type by currency
func ByCurrencyStr(currency string) (c currency, ok bool) {
	c, ok = currenciesByCurrency[currency]
	return
}

//ByNumberStr lookup for currency type by number
func ByNumberStr(number string) (c currency, ok bool) {
	c, ok = currenciesByNumber[number]
	return
}

//ByCountryStr lookup for currencies type by country
func ByCountryStr(country string) (c currencies, ok bool) {
	c, ok = currenciesByCountry[country]
	return
}

//ByCodeStrErr lookup for currency type by code
func ByCodeStrErr(code string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCode[code]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return
}

//ByCurrencyStrErr lookup for currency type by currency
func ByCurrencyStrErr(currencyStr string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 currency", currencyStr)
	}

	return
}

//ByNumberStrErr lookup for currency type by number
func ByNumberStrErr(number string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 number", number)
	}

	return
}

//ByCountryStrErr lookup for currencies type by country
func ByCountryStrErr(country string) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country]

	if !ok {
		return nil, fmt.Errorf("'%s' is not valid ISO-4217 country", country)
	}

	return
}

//ByCode lookup for currency type by code
func ByCode(code Code) (c currency, ok bool) {
	c, ok = currenciesByCode[code.String()]
	return
}

//ByCurrency lookup for currency type by currency
func ByCurrency(currency Currency) (c currency, ok bool) {
	c, ok = currenciesByCurrency[currency.String()]
	return
}

//ByNumber lookup for currency type by number
func ByNumber(number Number) (c currency, ok bool) {
	c, ok = currenciesByNumber[number.String()]
	return
}

//ByCountry lookup for currency type by country
func ByCountry(country Country) (c currencies, ok bool) {
	c, ok = currenciesByCountry[country.String()]
	return
}

//ByCodeErr lookup for currency type by code
func ByCodeErr(code Code) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCode[code.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return
}

//ByCurrencyErr lookup for currencies type by code
func ByCurrencyErr(currencyStr Currency) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 currency", currencyStr)
	}

	return
}

//ByNumberErr lookup for currencies type by number
func ByNumberErr(number Number) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 number", number)
	}

	return
}

//ByCountryErr lookup for currencies type by country
func ByCountryErr(country Country) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country.String()]

	if !ok {
		return nil, fmt.Errorf("'%s' is not valid ISO-4217 country", country)
	}

	return
}

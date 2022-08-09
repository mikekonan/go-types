/*
Package currency provides lookups methods over ISO-4217 currencies.

Constants have been generated with code generation ./generator.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package currency

import (
	"fmt"
)

type currency struct {
	countries     Countries
	currency      Currency
	code          Code
	number        Number
	decimalPlaces int
}

// Currency returns Currency
func (c currency) Currency() Currency { return c.currency }

// Code returns Code
func (c currency) Code() Code { return c.code }

// Number returns Number
func (c currency) Number() Number { return c.number }

// Countries returns Countries
func (c currency) Countries() Countries { return c.countries }

// DecimalPlaces returns DecimalPlaces
func (c currency) DecimalPlaces() int { return c.decimalPlaces }

type currencies []currency

// CurrencyByCurrency get currency by currency
func (currencies currencies) CurrencyByCurrency(curr string) (currency, bool) {
	for _, c := range currencies {
		if string(c.currency) == curr {
			return c, true
		}
	}

	return currency{}, false
}

// CurrencyByCode gets currency by code
func (currencies currencies) CurrencyByCode(code string) (currency, bool) {
	for _, c := range currencies {
		if string(c.code) == code {
			return c, true
		}
	}

	return currency{}, false
}

// CurrencyByNumber gets currency by number
func (currencies currencies) CurrencyByNumber(number string) (currency, bool) {
	for _, c := range currencies {
		if string(c.number) == number {
			return c, true
		}
	}

	return currency{}, false
}

// ByCodeStr lookup for currency type by code
func ByCodeStr(code string) (c currency, ok bool) {
	c, ok = currenciesByCode[code]
	return
}

// ByCurrencyStr lookup for currency type by currency
func ByCurrencyStr(currency string) (c currency, ok bool) {
	c, ok = currenciesByCurrency[currency]
	return
}

// ByNumberStr lookup for currency type by number
func ByNumberStr(number string) (c currency, ok bool) {
	c, ok = currenciesByNumber[number]
	return
}

// ByCountryStr lookup for currencies type by country
func ByCountryStr(country string) (c currencies, ok bool) {
	c, ok = currenciesByCountry[country]
	return
}

// ByCodeStrErr lookup for currency type by code
func ByCodeStrErr(code string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCode[code]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return
}

// ByCurrencyStrErr lookup for currency type by currency
func ByCurrencyStrErr(currencyStr string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 currency", currencyStr)
	}

	return
}

// ByNumberStrErr lookup for currency type by number
func ByNumberStrErr(number string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 number", number)
	}

	return
}

// ByCountryStrErr lookup for currencies type by country
func ByCountryStrErr(country string) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country]

	if !ok {
		return nil, fmt.Errorf("'%s' is not valid ISO-4217 country", country)
	}

	return
}

// ByCode lookup for currency type by code
func ByCode(code Code) (c currency, ok bool) {
	c, ok = currenciesByCode[code.String()]
	return
}

// ByCurrency lookup for currency type by currency
func ByCurrency(currency Currency) (c currency, ok bool) {
	c, ok = currenciesByCurrency[currency.String()]
	return
}

// ByNumber lookup for currency type by number
func ByNumber(number Number) (c currency, ok bool) {
	c, ok = currenciesByNumber[number.String()]
	return
}

// ByCountry lookup for currency type by country
func ByCountry(country Country) (c currencies, ok bool) {
	c, ok = currenciesByCountry[country.String()]
	return
}

// ByCodeErr lookup for currency type by code
func ByCodeErr(code Code) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCode[code.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 code", code)
	}

	return
}

// ByCurrencyErr lookup for currencies type by code
func ByCurrencyErr(currencyStr Currency) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 currency", currencyStr)
	}

	return
}

// ByNumberErr lookup for currencies type by number
func ByNumberErr(number Number) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number.String()]

	if !ok {
		return currency{}, fmt.Errorf("'%s' is not valid ISO-4217 number", number)
	}

	return
}

// ByCountryErr lookup for currencies type by country
func ByCountryErr(country Country) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country.String()]

	if !ok {
		return nil, fmt.Errorf("'%s' is not valid ISO-4217 country", country)
	}

	return
}

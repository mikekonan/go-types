/*
Package currency provides lookups methods over ISO-4217 currencies.

Constants have been generated with code generation ./generator.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package currency

import (
	"strings"
)

type currency struct {
	countries     Countries
	currency      Currency
	code          Code
	number        Number
	decimalPlaces int
	isHistorical  bool
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

// isHistorical returns isHistorical
func (c currency) IsHistorical() bool { return c.isHistorical }

type currencies []currency

// CurrencyByCurrency gets currency by currency
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
		if string(c.code) == strings.ToUpper(code) {
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
	c, ok = currenciesByCode[strings.ToUpper(code)]
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
	c, ok = currenciesByCode[strings.ToUpper(code)]

	if !ok {
		return currency{}, newInvalidDataError(code, standardISO4217Code)
	}

	return
}

// ByCurrencyStrErr lookup for currency type by currency
func ByCurrencyStrErr(currencyStr string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr]

	if !ok {
		return currency{}, newInvalidDataError(currencyStr, standardISO4217Currency)
	}

	return
}

// ByNumberStrErr lookup for currency type by number
func ByNumberStrErr(number string) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number]

	if !ok {
		return currency{}, newInvalidDataError(number, standardISO4217Number)
	}

	return
}

// ByCountryStrErr lookup for currencies type by country
func ByCountryStrErr(country string) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country]

	if !ok {
		return nil, newInvalidDataError(country, standardISO4217Country)
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
		return currency{}, newInvalidDataError(string(code), standardISO4217Code)
	}

	return
}

// ByCurrencyErr lookup for currencies type by code
func ByCurrencyErr(currencyStr Currency) (c currency, err error) {
	var ok bool
	c, ok = currenciesByCurrency[currencyStr.String()]

	if !ok {
		return currency{}, newInvalidDataError(string(currencyStr), standardISO4217Currency)
	}

	return
}

// ByNumberErr lookup for currencies type by number
func ByNumberErr(number Number) (c currency, err error) {
	var ok bool
	c, ok = currenciesByNumber[number.String()]

	if !ok {
		return currency{}, newInvalidDataError(string(number), standardISO4217Number)
	}

	return
}

// ByCountryErr lookup for currencies type by country
func ByCountryErr(country Country) (c currencies, err error) {
	var ok bool
	c, ok = currenciesByCountry[country.String()]

	if !ok {
		return nil, newInvalidDataError(string(country), standardISO4217Country)
	}

	return
}

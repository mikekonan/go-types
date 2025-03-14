package currency

import (
	"fmt"
)

type RegisteredCurrency struct {
	Currency      string
	Code          string
	Countries     Countries
	DecimalPlaces int
}

func RegisterCurrency(customCurrency RegisteredCurrency) error {
	if _, ok := currenciesByCode[customCurrency.Code]; ok {
		return fmt.Errorf("currency with code %s already registered", customCurrency.Code)
	}

	if _, ok := currenciesByCurrency[customCurrency.Currency]; ok {
		return fmt.Errorf("currency with name %s already registered", customCurrency.Currency)
	}

	var registeredCurrency = currency{
		countries:     customCurrency.Countries,
		currency:      Currency(customCurrency.Currency),
		code:          Code(customCurrency.Code),
		decimalPlaces: customCurrency.DecimalPlaces,
		isHistorical:  false,
	}

	currenciesByCode[customCurrency.Code] = registeredCurrency
	currenciesByCurrency[customCurrency.Currency] = registeredCurrency
	for _, country := range customCurrency.Countries {
		if _, ok := currenciesByCountry[string(country)]; ok {
			currenciesByCountry[string(country)] = append(currenciesByCountry[string(country)], registeredCurrency)
			continue
		}

		currenciesByCountry[string(country)] = []currency{registeredCurrency}
	}

	return nil
}

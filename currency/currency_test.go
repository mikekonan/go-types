package currency

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCurrenciesByCountry(t *testing.T) {
	for _, currencies := range currenciesByCountry {
		for _, currency := range currencies {
			//IsSet
			if !currency.Number().IsSet() {
				t.FailNow()
			}

			if !currency.Currency().IsSet() {
				t.FailNow()
			}

			if !currency.Code().IsSet() {
				t.FailNow()
			}

			for _, country := range currency.Countries() {
				if !country.IsSet() {
					t.FailNow()
				}
			}

			//IsCountryIn
			if !currency.countries.IsCountryIn(string(currency.countries[0])) {
				t.FailNow()
			}

			if currency.countries.IsCountryIn("") {
				t.FailNow()
			}

			//Stringer
			for _, country := range currency.Countries() {
				if len(country.String()) == 0 {
					t.FailNow()
				}
			}

			if len(currency.Currency().String()) == 0 {
				t.FailNow()
			}

			if len(currency.Code().String()) == 0 {
				t.FailNow()
			}

			if len(currency.Number().String()) == 0 {
				t.FailNow()
			}

			//driver.Valuer
			for _, country := range currency.Countries() {
				if _, err := country.Value(); err != nil {
					t.FailNow()
				}
			}

			if len(currency.Currency().String()) == 0 {
				t.FailNow()
			}

			if _, err := currency.Currency().Value(); err != nil {
				t.FailNow()
			}

			if _, err := currency.Code().Value(); err != nil {
				t.FailNow()
			}

			if _, err := currency.Number().Value(); err != nil {
				t.FailNow()
			}

			if _, err := Currency("").Value(); err != nil {
				t.FailNow()
			}

			if _, err := Currency("wrong").Value(); err == nil {
				t.FailNow()
			}

			if _, err := Code("").Value(); err != nil {
				t.FailNow()
			}

			if _, err := Code("wrong").Value(); err == nil {
				t.FailNow()
			}

			if _, err := Number("").Value(); err != nil {
				t.FailNow()
			}

			if _, err := Number("wrong").Value(); err == nil {
				t.FailNow()
			}

			if _, err := Country("").Value(); err != nil {
				t.FailNow()
			}

			if _, err := Country("wrong").Value(); err == nil {
				t.FailNow()
			}

			//json.Unmarshaler
			if err := json.Unmarshal([]byte(`""`), new(Code)); err == nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`""`), new(Country)); err == nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`""`), new(Number)); err == nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`""`), new(Currency)); err == nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`"AFN"`), new(Code)); err != nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`"BELGIUM"`), new(Country)); err != nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`"012"`), new(Number)); err != nil {
				t.FailNow()
			}

			if err := json.Unmarshal([]byte(`"US Dollar"`), new(Currency)); err != nil {
				t.FailNow()
			}
		}

		if _, ok := currencies.CurrencyByCode(""); ok {
			t.FailNow()
		}

		if _, ok := currencies.CurrencyByNumber(""); ok {
			t.FailNow()
		}

		if _, ok := currencies.CurrencyByCurrency(""); ok {
			t.FailNow()
		}

		if currency, ok := currencies.CurrencyByCode(currencies[0].Code().String()); !ok || !reflect.DeepEqual(currency, currencies[0]) {
			t.FailNow()
		}

		if currency, ok := currencies.CurrencyByNumber(currencies[0].Number().String()); !ok || !reflect.DeepEqual(currency, currencies[0]) {
			t.FailNow()
		}

		if currency, ok := currencies.CurrencyByCurrency(currencies[0].Currency().String()); !ok || !reflect.DeepEqual(currency, currencies[0]) {
			t.FailNow()
		}
	}
}

func TestCurrencyByCode(t *testing.T) {
	for _, currency := range currenciesByCode {
		//IsSet
		if !currency.Number().IsSet() {
			t.FailNow()
		}

		if !currency.Currency().IsSet() {
			t.FailNow()
		}

		if !currency.Code().IsSet() {
			t.FailNow()
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		//IsCountryIn
		if !currency.countries.IsCountryIn(string(currency.countries[0])) {
			t.FailNow()
		}

		if currency.countries.IsCountryIn("") {
			t.FailNow()
		}

		//Stringer
		for _, country := range currency.Countries() {
			if len(country.String()) == 0 {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Code().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Number().String()) == 0 {
			t.FailNow()
		}

		//driver.Valuer
		for _, country := range currency.Countries() {
			if _, err := country.Value(); err != nil {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if _, err := currency.Currency().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Code().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Number().Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Code("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Code("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Number("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Number("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Country("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Country("wrong").Value(); err == nil {
			t.FailNow()
		}

		//json.Unmarshaler
		if err := json.Unmarshal([]byte(`""`), new(Code)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Country)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Number)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Currency)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"AFN"`), new(Code)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"BELGIUM"`), new(Country)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"012"`), new(Number)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"US Dollar"`), new(Currency)); err != nil {
			t.FailNow()
		}
	}
}

func TestCurrencyByNumber(t *testing.T) {
	for _, currency := range currenciesByNumber {
		//IsSet
		if !currency.Number().IsSet() {
			t.FailNow()
		}

		if !currency.Currency().IsSet() {
			t.FailNow()
		}

		if !currency.Code().IsSet() {
			t.FailNow()
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		//IsCountryIn
		if !currency.countries.IsCountryIn(string(currency.countries[0])) {
			t.FailNow()
		}

		if currency.countries.IsCountryIn("") {
			t.FailNow()
		}

		//Stringer
		for _, country := range currency.Countries() {
			if len(country.String()) == 0 {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Code().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Number().String()) == 0 {
			t.FailNow()
		}

		//driver.Valuer
		for _, country := range currency.Countries() {
			if _, err := country.Value(); err != nil {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if _, err := currency.Currency().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Code().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Number().Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Code("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Code("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Number("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Number("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Country("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Country("wrong").Value(); err == nil {
			t.FailNow()
		}

		//json.Unmarshaler
		if err := json.Unmarshal([]byte(`""`), new(Code)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Country)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Number)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Currency)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"AFN"`), new(Code)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"BELGIUM"`), new(Country)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"012"`), new(Number)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"US Dollar"`), new(Currency)); err != nil {
			t.FailNow()
		}
	}
}

func TestCurrencyByByCurrency(t *testing.T) {
	for _, currency := range currenciesByCurrency {
		//IsSet
		if !currency.Number().IsSet() {
			t.FailNow()
		}

		if !currency.Currency().IsSet() {
			t.FailNow()
		}

		if !currency.Code().IsSet() {
			t.FailNow()
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		for _, country := range currency.Countries() {
			if !country.IsSet() {
				t.FailNow()
			}
		}

		//IsCountryIn
		if !currency.countries.IsCountryIn(string(currency.countries[0])) {
			t.FailNow()
		}

		if currency.countries.IsCountryIn("") {
			t.FailNow()
		}

		//Stringer
		for _, country := range currency.Countries() {
			if len(country.String()) == 0 {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Code().String()) == 0 {
			t.FailNow()
		}

		if len(currency.Number().String()) == 0 {
			t.FailNow()
		}

		//driver.Valuer
		for _, country := range currency.Countries() {
			if _, err := country.Value(); err != nil {
				t.FailNow()
			}
		}

		if len(currency.Currency().String()) == 0 {
			t.FailNow()
		}

		if _, err := currency.Currency().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Code().Value(); err != nil {
			t.FailNow()
		}

		if _, err := currency.Number().Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Currency("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Code("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Code("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Number("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Number("wrong").Value(); err == nil {
			t.FailNow()
		}

		if _, err := Country("").Value(); err != nil {
			t.FailNow()
		}

		if _, err := Country("wrong").Value(); err == nil {
			t.FailNow()
		}

		//json.Unmarshaler
		if err := json.Unmarshal([]byte(`""`), new(Code)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Country)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Number)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`""`), new(Currency)); err == nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"AFN"`), new(Code)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"BELGIUM"`), new(Country)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"012"`), new(Number)); err != nil {
			t.FailNow()
		}

		if err := json.Unmarshal([]byte(`"US Dollar"`), new(Currency)); err != nil {
			t.FailNow()
		}
	}
}

func TestByFunc(t *testing.T) {
	if _, ok := ByCodeStr("AFN"); !ok {
		t.FailNow()
	}

	if _, ok := ByCodeStr(""); ok {
		t.FailNow()
	}

	if _, ok := ByCurrencyStr("Afghani"); !ok {
		t.FailNow()
	}

	if _, ok := ByCurrencyStr(""); ok {
		t.FailNow()
	}

	if _, ok := ByNumberStr("012"); !ok {
		t.FailNow()
	}

	if _, ok := ByNumberStr(""); ok {
		t.FailNow()
	}

	if _, ok := ByCountryStr("BELGIUM"); !ok {
		t.FailNow()
	}

	if _, ok := ByCountryStr(""); ok {
		t.FailNow()
	}

	if _, err := ByCodeStrErr("AFN"); err != nil {
		t.FailNow()
	}

	if _, err := ByCodeStrErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByCurrencyStrErr("Afghani"); err != nil {
		t.FailNow()
	}

	if _, err := ByCurrencyStrErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByNumberStrErr("012"); err != nil {
		t.FailNow()
	}

	if _, err := ByNumberStrErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByCountryStrErr("BELGIUM"); err != nil {
		t.FailNow()
	}

	if _, err := ByCountryStrErr(""); err == nil {
		t.FailNow()
	}

	if _, ok := ByCode("AFN"); !ok {
		t.FailNow()
	}

	if _, ok := ByCode(""); ok {
		t.FailNow()
	}

	if _, ok := ByCurrency("Afghani"); !ok {
		t.FailNow()
	}

	if _, ok := ByCurrency(""); ok {
		t.FailNow()
	}

	if _, ok := ByNumber("012"); !ok {
		t.FailNow()
	}

	if _, ok := ByNumber(""); ok {
		t.FailNow()
	}

	if _, ok := ByCountry("BELGIUM"); !ok {
		t.FailNow()
	}

	if _, ok := ByCountry(""); ok {
		t.FailNow()
	}

	if _, err := ByCodeErr("AFN"); err != nil {
		t.FailNow()
	}

	if _, err := ByCodeErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByCurrencyErr("Afghani"); err != nil {
		t.FailNow()
	}

	if _, err := ByCurrencyErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByNumberErr("012"); err != nil {
		t.FailNow()
	}

	if _, err := ByNumberErr(""); err == nil {
		t.FailNow()
	}

	if _, err := ByCountryErr("BELGIUM"); err != nil {
		t.FailNow()
	}

	if _, err := ByCountryErr(""); err == nil {
		t.FailNow()
	}
}

func TestCurrencies(t *testing.T) {
	c, _ := ByCurrencyStr("Euro")
	countries := c.Countries()

	countries.IsCountryIn("FINLAND")
}

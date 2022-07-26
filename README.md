[![Go Report Card](https://goreportcard.com/badge/github.com/mikekonan/go-types)](https://goreportcard.com/report/github.com/mikekonan/go-types) [![Build Status](https://travis-ci.com/mikekonan/go-types.svg?branch=main)](https://travis-ci.com/mikekonan/go-types) [![codecov](https://codecov.io/gh/mikekonan/go-types/branch/main/graph/badge.svg?token=83Q04OW4I1)](https://codecov.io/gh/mikekonan/go-types)
# go-types
This library has been created with the purpose to facilitate the store, validation, and transfer of Go ISO-3166/ISO-4217/timezones/emails/URL types. There is a [openapi3 spec](https://github.com/mikekonan/go-types/blob/main/swagger.yaml) of that type and make you able to include it into your spec. All types has own ozzo.Validate, json.Unmarshaler, Stringer and driver.Valuer implementations.

# Installation
```bash
go get github.com/mikekonan/go-types/v2
```
# Usage:
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mikekonan/go-types/v2/country"
	"github.com/mikekonan/go-types/v2/country/alpha2"
	"github.com/mikekonan/go-types/v2/country/alpha3"
	"github.com/mikekonan/go-types/v2/country/name"
	"github.com/mikekonan/go-types/v2/currency"
	"github.com/mikekonan/go-types/v2/currency/code"
)

// 1. use in your structs
type User struct {
	Name     string             `json:"name" db:"name"`
	Country  country.Alpha2Code `json:"country" db:"country"`
	Currency currency.Code      `json:"currency" db:"currency"`
}

func main() {
	// 2. use in your wire
	user := User{}
	_ = json.Unmarshal([]byte(`{"name":"name", "country": "CA", "currency": "CAD"}`), &user)

	// 3. check is set
	user.Country.IsSet()
	user.Currency.IsSet()

	// 4. validate using ozzo-validation
	if err := validation.ValidateStruct(&user, validation.Field(&user.Country), validation.Field(&user.Currency)); err != nil {
		log.Fatal(err)
	}

	// 5. lookup by alpha2, alpha3, country name
	if userCountry, ok := country.ByAlpha2Code(user.Country); ok {
		fmt.Printf("country name - '%s', alpha-2 - '%s', alpha-3 - '%s'", userCountry.Name(), userCountry.Alpha2Code(), userCountry.Alpha3Code())
	}

	// 5. lookup by currency code
	if userCurrency, ok := currency.ByCode(user.Currency); ok {
		fmt.Printf("currency name - '%s', code - '%s', number - '%s', countries - '%s', decimal places - '%d'",
			userCurrency.Currency(), userCurrency.Code(), userCurrency.Number(), userCurrency.Countries(), userCurrency.DecimalPlaces())
	}

	// 6. store in db
	fmt.Println(user.Country.Value())  //prints 'CA'
	fmt.Println(user.Currency.Value()) //prints 'CAN'

	// 7. use specific country constants
	fmt.Println(country.Canada.Alpha2Code())
	fmt.Println("name:", name.Canada)
	fmt.Println("alpha-2:", alpha2.CA)
	fmt.Println("alpha-3:", alpha3.CAN)

	// 8. use specific currency codes
	fmt.Println(code.CAD)
}
```

## Links:
- Currency Codes [(ISO 4217)](https://www.currency-iso.org/en/home/tables/table-a1.html) - [wiki](https://en.wikipedia.org/wiki/ISO_4217)
- Country Codes [(ISO 3166)](https://www.iso.org/iso-3166-country-codes.html) - [wiki](https://en.wikipedia.org/wiki/ISO_3166-2)
- URL(including HttpURL) [(standard)](https://url.spec.whatwg.org/) - [wiki](https://en.wikipedia.org/wiki/URL)
- Email [(part of RFC5322)](https://tools.ietf.org/html/rfc5322) - [wiki](https://en.wikipedia.org/wiki/Email_address)
- Timezone [(RFC6557 IANA timezones)](https://www.iana.org/time-zones) - [wiki](https://en.wikipedia.org/wiki/Time_zone)

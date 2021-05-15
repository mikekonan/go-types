[![Go Report Card](https://goreportcard.com/badge/github.com/mikekonan/go-types)](https://goreportcard.com/report/github.com/mikekonan/go-types) [![Build Status](https://travis-ci.com/mikekonan/go-types.svg?branch=main)](https://travis-ci.com/mikekonan/go-types) [![codecov](https://codecov.io/gh/mikekonan/go-types/branch/main/graph/badge.svg?token=83Q04OW4I1)](https://codecov.io/gh/mikekonan/go-types)
# go-types
This library has been created with the purpose to facilitate the store, validation, and transfer of Go ISO-3166/ISO-4217/emails/URL types. There is a [openapi3 spec](https://github.com/mikekonan/go-types/blob/main/swagger.yaml) of that type and make you able to include it into your spec. All types has own ozzo.Validate, json.Unmarshaler, Stringer and driver.Valuer implementations.

# Installation
```bash
go get github.com/mikekonan/go-types
```
# Usage:
```go
//1. use in your structs
type User struct {
	Name    string             `json:"name" db:"name"`
	Country country.Alpha2Code `json:"country" db:"country"`
}

func main() {
	user := User{}
	//2. use in your wire
	json.Unmarshal([]byte(`{"name":"name", "country": "ca"}`), &user)
	//3. check is set
	user.Country.IsSet() //check user country is provided
	//4. validate using ozzo-validation
	if err := validation.ValidateStruct(&user, validation.Field(&user.ountry, validation.Required, user.Country)); err != nil {
		log.Fatal(err)
	}
	//5. lookup by alpha2, alpha3, country name
	if userCountry, ok := country.ByAlpha2Code(user.Country); ok {
		fmt.Printf("country name - '%s', alpha-2 - '%s', alpha-3 - '%s'", serCountry.Name(), userCountry.Alpha2Code(), userCountry.lpha3Code())
	}
	//6. store in db
	fmt.Println(user.Country.Value()) //prints 'CA'
	//7. use specific countries
	fmt.Println(country.Canada.Alpha2Code())
}
```

## Links:
- Currency Codes [(ISO 4217)](https://www.currency-iso.org/en/home/tables/table-a1.html) - [wiki](https://en.wikipedia.org/wiki/ISO_4217)
- Country Codes [(ISO 3166)](https://www.iso.org/iso-3166-country-codes.html) - [wiki](https://en.wikipedia.org/wiki/ISO_3166-2)
- URL [(standard)](https://url.spec.whatwg.org/) - [wiki](https://en.wikipedia.org/wiki/URL)
- Email [(part of RFC5322)](https://tools.ietf.org/html/rfc5322) - [wiki](https://en.wikipedia.org/wiki/Email_address)

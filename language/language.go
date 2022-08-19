package language

import (
	"fmt"
	"strings"
)

// Language struct represents language entity according to ISO 639-1/693-2 standards.
type Language struct {
	name   Name
	alpha2 Alpha2Code
	alpha3 Alpha3Code
}

// Name returns language name
func (language Language) Name() Name { return language.name }

// Alpha2Code returns alpha-2 code
func (language Language) Alpha2Code() Alpha2Code { return language.alpha2 }

// Alpha3Code returns alpha-2 code
func (language Language) Alpha3Code() Alpha3Code { return language.alpha3 }

// NameStr returns language language string
func (language Language) NameStr() string { return language.name.String() }

// Alpha2CodeStr returns language alpha-2 code string
func (language Language) Alpha2CodeStr() string { return language.alpha2.String() }

// Alpha3CodeStr returns language alpha-2 code string
func (language Language) Alpha3CodeStr() string { return language.alpha3.String() }

// ByAlpha3Code lookup for language by alpha-3 code
func ByAlpha3Code(code Alpha3Code) (result Language, ok bool) {
	result, ok = LanguageByAlpha3[strings.ToLower(code.String())]

	return
}

// ByAlpha3CodeStr lookup for language by alpha-3 code string
func ByAlpha3CodeStr(code string) (result Language, ok bool) {
	return ByAlpha3Code(Alpha3Code(code))
}

// ByAlpha3CodeErr lookup for language by alpha-3 code with error return type
func ByAlpha3CodeErr(code Alpha3Code) (result Language, err error) {
	var ok bool
	result, ok = ByAlpha3Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO 639-2 code", code)
	}

	return
}

// ByAlpha3CodeStrErr lookup for language by alpha-3 code string with error return type
func ByAlpha3CodeStrErr(code string) (result Language, err error) {
	return ByAlpha3CodeErr(Alpha3Code(code))
}

// ByAlpha2Code lookup for language by alpha-2 code
func ByAlpha2Code(code Alpha2Code) (result Language, ok bool) {
	result, ok = LanguageByAlpha2[strings.ToLower(code.String())]
	return
}

// ByAlpha2CodeStr lookup for language by alpha-2 code string
func ByAlpha2CodeStr(code string) (result Language, ok bool) {
	return ByAlpha2Code(Alpha2Code(code))
}

// ByAlpha2CodeErr lookup for language by alpha-2 code with error return type
func ByAlpha2CodeErr(code Alpha2Code) (result Language, err error) {
	var ok bool
	result, ok = ByAlpha2Code(code)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO 639-1 code", code)
	}

	return
}

// ByAlpha2CodeStrErr lookup for language by alpha-2 code string with error return type
func ByAlpha2CodeStrErr(code string) (Language, error) {
	return ByAlpha2CodeErr(Alpha2Code(code))
}

// ByName lookup for language by name
func ByName(language Name) (result Language, ok bool) {
	result, ok = LanguageByName[language.String()]
	return
}

// ByNameStr lookup for language by name string
func ByNameStr(language string) (Language, bool) {
	return ByName(Name(language))
}

// ByNameErr lookup for language by name with error return type
func ByNameErr(language Name) (result Language, err error) {
	var ok bool
	result, ok = ByName(language)
	if !ok {
		err = fmt.Errorf("'%s' is not valid ISO 639-1 Language name", language)
	}

	return
}

// ByNameStrErr lookup for language by name string with error return type
func ByNameStrErr(language string) (Language, error) {
	return ByNameErr(Name(language))
}

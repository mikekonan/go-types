package bcp47_language

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/mikekonan/go-types/v2/language"
	stdLanguage "golang.org/x/text/language"
)

// Language represents a BCP49 language entity.
// It includes the base ISO639 language, the original BCP47 code string,
// and the corresponding language.Tag from the golang.org/x/text/language package.
type Language struct {
	raw    string
	rawTag stdLanguage.Tag
}

// ByStrErr creates a new Language by parsing a BCP47 code string.
// It returns an error if the BCP47 code is invalid or if the base ISO639 language cannot be parsed.
func ByStrErr(code string) (result Language, err error) {
	if "" == code {
		return Language{}, fmt.Errorf("code is empty string")
	}

	result.rawTag, err = stdLanguage.Parse(code)
	if err != nil {
		return Language{}, fmt.Errorf("'%s' is not valid BCP47 language code", code)
	}

	result.raw = code

	return result, nil
}

// BaseISO639Language returns the base ISO639 language of the BCP47 code.
// The returned language.Language represents the most significant part of the language
// identification used in the BCP47 code. For example, for the BCP47 code "en-US",
// the base ISO639 language would be "en".
//
// This method will return an error if the ISO639 language cannot be determined from the BCP47 code.
func (l Language) BaseISO639Language() (language.Language, error) {
	baseRaw, _ := l.rawTag.Base()
	baseISO639Language, err := language.ByAlpha2CodeStrErr(baseRaw.String())
	if err != nil {
		return language.Language{}, err
	}

	return baseISO639Language, nil
}

// String returns the original BCP47 code string.
func (l Language) String() string {
	return l.raw
}

// Raw returns the parsed BCP47 code as a language.Tag.
func (l Language) Raw() stdLanguage.Tag {
	return l.rawTag
}

// UnmarshalJSON unmarshals a JSON-encoded BCP47 code into a Language.
func (l *Language) UnmarshalJSON(data []byte) (err error) {
	var str string
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	entity, err := ByStrErr(str)
	if err != nil {
		return err
	}

	*l = entity
	return nil
}

// Value returns the BCP47 code as a driver.Value for use with SQL databases.
func (l *Language) Value() (value driver.Value, err error) {
	return l.raw, nil
}

// Validate checks if the Language is valid by attempting to parse its BCP47 code.
// It returns an error if the BCP47 code is invalid.
func (l Language) Validate() (err error) {
	_, err = ByStrErr(l.raw)

	return
}

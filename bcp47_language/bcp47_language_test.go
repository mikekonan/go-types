package bcp47_language

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"reflect"
=======
>>>>>>> main
	"testing"

	"github.com/mikekonan/go-types/v2/language"
)

var testCases = []struct {
	name                      string
	input                     []byte
	bca47ParseExpectedError   string
	iso639LookupExpectedError string
	jsonUnmarshalError        string
	iso639expectedLang        language.Language
	isBCA49Valid              bool
}{
	{
		name:                    "Valid language code no region",
		input:                   []byte(`{"language":"fr"}`),
		bca47ParseExpectedError: "",
		iso639expectedLang:      language.French,
		isBCA49Valid:            true,
	},
	{
		name:                    "Valid language code",
		input:                   []byte(`{"language":"en-US"}`),
		bca47ParseExpectedError: "",
		iso639expectedLang:      language.English,
		isBCA49Valid:            true,
	},
	{
		name:                      "Local language code",
		input:                     []byte(`{"language":"qaa-Qaaa-QM-x-southern"}`),
		bca47ParseExpectedError:   "",
		isBCA49Valid:              true,
		iso639LookupExpectedError: "'qaa' is not valid ISO 639-1 code",
	},
	{
		name:               "Invalid language code",
		input:              []byte(`{"language":"invalid"}`),
		jsonUnmarshalError: "'invalid' is not valid BCP47 language code",
	},
	{
		name:               "Empty language code",
		input:              []byte(`{"language":""}`),
		jsonUnmarshalError: "code is empty string",
	},
	{
		name:                      "Non-standard language code",
		input:                     []byte(`{"language":"i-klingon"}`),
		isBCA49Valid:              true,
		iso639LookupExpectedError: "'tlh' is not valid ISO 639-1 code",
	},
}

type languageStruct struct {
	Language Language `json:"language"`
}

func TestLanguage(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lStruct := new(languageStruct)
			if err := json.Unmarshal(tc.input, lStruct); err != nil {
				if err == nil && tc.jsonUnmarshalError != "" {
					t.Errorf("UnmarshalJSON() with input %s: expected error %s, but got %s", tc.input, tc.jsonUnmarshalError, err)
					return
				}

				if err != nil {
					if err.Error() != tc.jsonUnmarshalError {
						t.Errorf("UnmarshalJSON() with input %s: expected error %s, but got %s", tc.input, tc.jsonUnmarshalError, err)
					}

					return
				}

				return
			}

<<<<<<< HEAD
			data, err := json.Marshal(lStruct)
			if err != nil {
				t.Errorf("MarshalJSON() with input %s: expecting no error, but got %s", tc.input, err)
				return
			}

			if !reflect.DeepEqual(data, tc.input) {
				t.Errorf("MarshalJSON() with input %s not equals to %s", tc.input, data)
				return
			}

=======
>>>>>>> main
			if tc.isBCA49Valid && lStruct.Language.Validate() != nil {
				t.Errorf("Validate() with input %s: expected to be valid %v", tc.input, tc.isBCA49Valid)
				return
			}

			value, err := lStruct.Language.Value()
			if value != lStruct.Language.String() || err != nil {
				t.Errorf("Value() with input %s: expected to be equal to %s without error, but %s", lStruct.Language.String(), value, err)
				return
			}

			if lStruct.Language.rawTag != lStruct.Language.Raw() {
				t.Errorf("Raw() with input %s: expected to be equal to %s without error, but %s", lStruct.Language.String(), lStruct.Language.rawTag, err)
				return
			}

			lang, err := lStruct.Language.BaseISO639Language()
			if err != nil {
				if tc.iso639LookupExpectedError == "" || err.Error() != tc.iso639LookupExpectedError {
					t.Errorf("BaseISO639Language() with input %s: expected error %s, but got %s", tc.input, tc.iso639LookupExpectedError, err)
					return
				}
			}

			if tc.iso639expectedLang != lang {
				t.Errorf("BaseISO639Language() with input %s: expected lang %s, but got %s", tc.input, tc.iso639expectedLang.Alpha2CodeStr(), lang.Alpha2CodeStr())
			}
		})
	}
}
<<<<<<< HEAD

func TestLanguage2(t *testing.T) {
	lang, _ := ByStrErr("en-US")
	l := &languageStruct{Language: lang}
	data, _ := json.Marshal(l)
	fmt.Println(data)
}
=======
>>>>>>> main

package phone

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var testDialCodeCases = []struct {
	testValue              DialCode
	expectingValidateError bool
	expectingValueError    bool
}{
	{testValue: "", expectingValidateError: true, expectingValueError: false},
	{testValue: "12", expectingValidateError: true, expectingValueError: true},
	{testValue: "1252365", expectingValidateError: true, expectingValueError: true},
	{testValue: "rqweq", expectingValidateError: true, expectingValueError: true},
	{testValue: "r12", expectingValidateError: true, expectingValueError: true},
	{testValue: "999", expectingValidateError: true, expectingValueError: true},
}

func TestDialCode_Validate(t *testing.T) {
	for _, testCase := range testDialCodeCases {
		actualErr := testCase.testValue.Validate()
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestDialCode_UnmarshalJSON(t *testing.T) {
	for _, testCase := range testDialCodeCases {
		jsonStr := fmt.Sprintf("%s", testCase.testValue)
		var dialCode DialCode

		actualErr := json.Unmarshal([]byte(jsonStr), &dialCode)
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestDialCode_Value(t *testing.T) {
	for _, testCase := range testDialCodeCases {
		actualValue, actualErr := testCase.testValue.Value()
		if (actualErr == nil) == testCase.expectingValueError {
			t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValueError)
		}

		if ((actualErr == nil) == testCase.expectingValueError) && actualValue != string(testCase.testValue) {
			t.Errorf(`Value: '%s'. value differs: expect: '%s', actual: '%s'"`, testCase.testValue, testCase.testValue, actualValue)
		}
	}
}

func TestMappingIsCorrect(t *testing.T) {
	for key, language := range LanguageByName {
		if Name(key) != language.name {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha3 {
		if Alpha3Code(key) != language.alpha3 {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha2 {
		if Alpha2Code(key) != language.alpha2 {
			t.FailNow()
		}
	}
}

func TestMappingStringsCorrect(t *testing.T) {
	for key, language := range LanguageByName {
		if Name(key).String() != language.name.String() {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha3 {
		if Alpha3Code(key).String() != language.alpha3.String() {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha2 {
		if Alpha2Code(key).String() != language.alpha2.String() {
			t.FailNow()
		}
	}
}

func TestAlpha2Lookup(t *testing.T) {
	if _, ok := ByAlpha2Code("l"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr("l"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr("l"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr("l"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha2Code(Latvian.Alpha2Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr(strings.ToUpper(Latvian.Alpha2Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr(Latvian.Alpha2Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr(strings.ToUpper(Latvian.Alpha2Code().String())); err != nil {
		t.FailNow()
	}
}

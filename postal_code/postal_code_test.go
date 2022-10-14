package postalcode

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testCases = []struct {
	testValue               PostalCode
	expectingValidateError  bool
	expectingValueError     bool
	expectingUnmarshalError bool
}{
	{testValue: "", expectingValidateError: true, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "CA1-4124", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "CA12 4124", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: " CA12 4122344", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "{}", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "112#124", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "sdfsdf#asd", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "a2g0)d23", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "12312312313221", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "CA124124", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "YYAAAAAA", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "1234567", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "\aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com",
		expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
}

func TestPostalCode_Validate(t *testing.T) {
	for _, testCase := range testCases {
		actualErr := testCase.testValue.Validate()
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestPostalCode_UnmarshalJSON(t *testing.T) {
	for _, testCase := range testCases {
		jsonStr := fmt.Sprintf(`{"postalCode":"%s"}`, testCase.testValue)
		var PostalCodeStruct struct {
			PostalCode PostalCode `json:"postalCode"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &PostalCodeStruct)
		if (actualErr == nil) == testCase.expectingUnmarshalError {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestPostalCode_Value(t *testing.T) {
	for _, testCase := range testCases {
		actualValue, actualErr := testCase.testValue.Value()
		if (actualErr == nil) == testCase.expectingValueError {
			t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValueError)
		}

		if ((actualErr == nil) == testCase.expectingValueError) && actualValue != string(testCase.testValue) {
			t.Errorf(`Value: '%s'. value differs: expect: '%s', actual: '%s'"`, testCase.testValue, testCase.testValue, actualValue)
		}
	}
}

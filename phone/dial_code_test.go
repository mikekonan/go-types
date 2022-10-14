package phone

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testDialCodeCases = []struct {
	testValue               DialCode
	expectingValidateError  bool
	expectingValueError     bool
	expectingUnmarshalError bool
}{
	{testValue: "", expectingValidateError: true, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "12", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "1252365", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: " 1252365", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "rqweq", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "r12", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "999", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "1", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
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
		jsonStr := fmt.Sprintf(`{"dialCode":"%s"}`, testCase.testValue)
		var dialCodeStruct struct {
			DialCode DialCode `json:"dialCode"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &dialCodeStruct)
		if (actualErr == nil) == testCase.expectingUnmarshalError {
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

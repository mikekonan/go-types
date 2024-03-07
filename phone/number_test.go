package phone

import (
	"encoding/json"
	"fmt"
	"github.com/mikekonan/go-types/v2/country"
	"testing"
)

var testNumberCases = []struct {
	testValue               Number
	expectingValidateError  bool
	expectingValueError     bool
	expectingUnmarshalError bool
}{
	{testValue: "", expectingValidateError: true, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "123", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "12345", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "1234", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "1234567", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "1234567890", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "1234567890123", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "123456789012345", expectingValidateError: false, expectingValueError: false, expectingUnmarshalError: false},
	{testValue: "1234567888888888", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "sss", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "qehweter", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "kslhdofighsudi", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "ewf123", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "sdrgte1241", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "email@example.com", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
	{testValue: "email@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com",
		expectingValidateError: true, expectingValueError: true, expectingUnmarshalError: true},
}

func TestNumber_Validate(t *testing.T) {
	for _, testCase := range testNumberCases {
		actualErr := testCase.testValue.Validate()
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestNumber_UnmarshalJSON(t *testing.T) {
	for _, testCase := range testNumberCases {
		jsonStr := fmt.Sprintf(`{"phoneNumber":"%s"}`, testCase.testValue)
		var phoneNumberStruct struct {
			Number Number `json:"phoneNumber"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &phoneNumberStruct)
		if (actualErr == nil) == testCase.expectingUnmarshalError {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValidateError)
		}
	}
}

func TestNumber_Value(t *testing.T) {
	for _, testCase := range testNumberCases {
		actualValue, actualErr := testCase.testValue.Value()
		if (actualErr == nil) == testCase.expectingValueError {
			t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.testValue, testCase.expectingValueError)
		}

		if ((actualErr == nil) == testCase.expectingValueError) && actualValue != string(testCase.testValue) {
			t.Errorf(`Value: '%s'. value differs: expect: '%s', actual: '%s'"`, testCase.testValue, testCase.testValue, actualValue)
		}
	}
}

var parseTestCases = []struct {
	testPhoneNumber string
	testCountryCode country.Alpha2Code
	expectingError  bool
}{
	{testPhoneNumber: "", testCountryCode: country.Australia.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "sdfge5uy2fwdf", testCountryCode: country.Australia.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "123", testCountryCode: country.Australia.Alpha2Code(), expectingError: false},
	{testPhoneNumber: "375295555555", testCountryCode: country.Belgium.Alpha2Code(), expectingError: false},
	{testPhoneNumber: "375295555555", testCountryCode: country.Belarus.Alpha2Code(), expectingError: false},
}

func TestParse(t *testing.T) {
	for _, testCase := range parseTestCases {
		_, _, actualErr := Parse(testCase.testPhoneNumber, testCase.testCountryCode)
		if (actualErr == nil) && testCase.expectingError {
			t.Errorf(`PhoneNumber: '%s', CountryCode: '%s'. Error is expected but the result was opposite`, testCase.testPhoneNumber, testCase.testCountryCode)
		}
	}
}

var strictParseTestCases = []struct {
	testPhoneNumber string
	testCountryCode country.Alpha2Code
	expectingError  bool
}{
	{testPhoneNumber: "", testCountryCode: country.Australia.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "asdgdhf", testCountryCode: country.Australia.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "123", testCountryCode: country.Australia.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "375295555555", testCountryCode: country.Belgium.Alpha2Code(), expectingError: true},
	{testPhoneNumber: "375295555555", testCountryCode: country.Belarus.Alpha2Code(), expectingError: false},
}

func TestParseWithRegionCheck(t *testing.T) {
	for _, testCase := range strictParseTestCases {
		_, _, actualErr := StrictParse(testCase.testPhoneNumber, testCase.testCountryCode)
		if (actualErr == nil) && testCase.expectingError {
			t.Errorf(`PhoneNumber: '%s', CountryCode: '%s'. Error is expected but the result was opposite`, testCase.testPhoneNumber, testCase.testCountryCode)
		}
	}
}

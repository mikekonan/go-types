package email

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testCases = []struct {
	emailUnderTest         Email
	expectingValidateError bool
	expectingValueError    bool
}{
	{emailUnderTest: "", expectingValidateError: true, expectingValueError: false},
	{emailUnderTest: "@", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: " @", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: " @ ", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "@ ", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "@.", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "..@.", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: ".@..", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "@@..", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: ".@.", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: ".@", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@x", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: ".email@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email.@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email..test@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: ".email..test.@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@at@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "some whitespace@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@whitespace example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", expectingValidateError: true, expectingValueError: true},
	{emailUnderTest: "email@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com",
		expectingValidateError: true, expectingValueError: true},

	{emailUnderTest: "email@gmail.com", expectingValidateError: false, expectingValueError: false},
	{emailUnderTest: "email.email@gmail.com", expectingValidateError: false, expectingValueError: false},
	{emailUnderTest: "email_email@gmail.com", expectingValidateError: false, expectingValueError: false},
	{emailUnderTest: "email+extra@example.com", expectingValidateError: false, expectingValueError: false},
	{emailUnderTest: "EMAIL@aol.co.uk", expectingValidateError: false, expectingValueError: false},
	{emailUnderTest: "EMAIL+EXTRA@aol.co.uk", expectingValidateError: false, expectingValueError: false},
}

func TestEmailValidate(t *testing.T) {
	for _, testCase := range testCases {
		actualErr := testCase.emailUnderTest.Validate()
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.emailUnderTest, testCase.expectingValidateError)
		}
	}
}

func TestEmailJsonUnmarshal(t *testing.T) {
	for _, testCase := range testCases {
		jsonStr := fmt.Sprintf(`{"email":"%s"}`, testCase.emailUnderTest)
		var emailStruct struct {
			Email Email `json:"email"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &emailStruct)
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.emailUnderTest, testCase.expectingValidateError)
		}
	}
}

func TestEmailValue(t *testing.T) {
	for _, testCase := range testCases {
		actualValue, actualErr := testCase.emailUnderTest.Value()
		if (actualErr == nil) == testCase.expectingValueError {
			t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.emailUnderTest, testCase.expectingValueError)
		}

		if ((actualErr == nil) == testCase.expectingValueError) && actualValue != string(testCase.emailUnderTest) {
			t.Errorf(`Value: '%s'. value differs: expect: '%s', actual: '%s'"`, testCase.emailUnderTest, testCase.emailUnderTest, actualValue)
		}
	}
}

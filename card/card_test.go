package card

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var testCases = []struct {
	cardDate    CardDate
	isExpectErr bool
}{
	// err
	{"", true},
	{"00/00", true},
	{"1", true},
	{"as/s", true},
	{"01", true},
	{"0221", true},
	{"01\\21", true},
	{"01//21", true},
	{"null", true},
	{"nil", true},
	{"01/2021", true},
	{"13/21", true},
	{"year", true},

	// no err
	{"01/21", false},
	{"12/21", false},
	{"01/99", false},
}

func TestCardDateJsonUnmarshal(t *testing.T) {
	for _, testCase := range testCases {
		jsonStr := fmt.Sprintf(`{"cardDate":"%s"}`, testCase.cardDate)
		var cardStruct struct {
			CardDate CardDate `json:"cardDate"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &cardStruct)
		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.cardDate, testCase.isExpectErr)
		}
	}
}

var testCasesValidation = []struct {
	cardDate    CardDate
	isExpectErr bool
}{
	// err
	{"", true},
	{"00/00", true},
	{"1", true},
	{"as/s", true},
	{"01", true},
	{"0221", true},
	{"01\\21", true},
	{"01//21", true},
	{"null", true},
	{"nil", true},
	{"01/2021", true},
	{"01/68", false},
	{"01/69", false},
	{"01/99", false},
	{"13/21", true},
	{"year", true},
	{CardDate(time.Now().AddDate(-1, 0, 0).Format(layout)), true},
	{CardDate(time.Now().AddDate(0, -1, 0).Format(layout)), true},

	// no err
	{CardDate(time.Now().Format(layout)), false},
	{CardDate(time.Now().AddDate(1, 0, 0).Format(layout)), false},
	{CardDate(time.Now().AddDate(0, 1, 0).Format(layout)), false},
}

func TestCardDateValue(t *testing.T) {
	for _, testCase := range testCasesValidation {
		actualVal, actualErr := testCase.cardDate.Value()

		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.cardDate, testCase.isExpectErr)
		}

		if testCase.isExpectErr && actualVal != nil {
			t.Errorf("Value should be empty")
		}
	}
}

func TestCardDateValidate(t *testing.T) {
	for _, testCase := range testCasesValidation {
		actualErr := testCase.cardDate.Validate()

		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.cardDate, testCase.isExpectErr)
		}
	}
}

package time

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"
)

type dt struct {
	XMLName xml.Name `json:"-" xml:"test"`
	D       Date     `json:"d,omitempty" xml:"d,omitempty"`
}

type unmarshalType struct {
	testJSON               []byte
	testXML                []byte
	expectingValidateError bool
}

var unmarshalTestCases = []unmarshalType{
	{
		testJSON:               []byte(`{}`),
		testXML:                []byte(`<?xml version="1.0"?><test></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"d": ""}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d></d></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"d": "2006-01-02"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d>2006-01-02</d></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"d": 1234}`),
		testXML:                []byte(`test><d>@</d>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"d": "@"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d>@</d></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"d": "2006-01-02T00:00:00"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d>2006-01-02T00:00:00</d></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"d": "00:00:00"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d>00:00:00</d></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"d": "text-text-text"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><d>text-text-text</d></test>`),
		expectingValidateError: true,
	},
}

func TestUnmarshalJSON(t *testing.T) {
	testUnmarshal(t, json.Unmarshal, func(u unmarshalType) []byte {
		return u.testJSON
	})
}

func TestUnmarshalXML(t *testing.T) {
	testUnmarshal(t, xml.Unmarshal, func(u unmarshalType) []byte {
		return u.testXML
	})
}

func testUnmarshal(t *testing.T, unmarshalFunc func([]byte, any) error, valueFunc func(unmarshalType) []byte) {
	for _, testCase := range unmarshalTestCases {
		var d dt
		actualErr := unmarshalFunc(valueFunc(testCase), &d)
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, valueFunc(testCase), testCase.expectingValidateError)
		}
	}
}

type marshalType struct {
	testData     dt
	expectedJSON []byte
	expectedXML  []byte
}

var marshalTestCases = []marshalType{
	{
		testData:     dt{},
		expectedJSON: []byte(`{"d":null}`),
		expectedXML:  []byte(`<test></test>`),
	},
	{
		testData:     dt{D: DateFromTime(time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))},
		expectedJSON: []byte(`{"d":"2006-01-02"}`),
		expectedXML:  []byte(`<test><d>2006-01-02</d></test>`),
	},
	{
		testData:     dt{D: DateFromTime(time.Date(1996, 8, 30, 10, 11, 40, 0, time.UTC))},
		expectedJSON: []byte(`{"d":"1996-08-30"}`),
		expectedXML:  []byte(`<test><d>1996-08-30</d></test>`),
	},
}

func TestMarshalJSON(t *testing.T) {
	testMarshal(t, "json", json.Marshal, func(m marshalType) []byte {
		return m.expectedJSON
	})
}

func TestMarshalXML(t *testing.T) {
	testMarshal(t, "xml", xml.Marshal, func(m marshalType) []byte {
		return m.expectedXML
	})
}

func testMarshal(t *testing.T, tp string, marshalFunc func(any) ([]byte, error), valueFunc func(marshalType) []byte) {
	for _, testCase := range marshalTestCases {
		actual, actualErr := marshalFunc(testCase.testData)
		if actualErr != nil {
			t.Errorf(`%s`, actualErr)
		}
		if string(actual) != string(valueFunc(testCase)) {
			t.Errorf("Actual %s does not equal expected. Actual: '%s', expected: '%s'", tp, actual, valueFunc(testCase))
		}
	}
}

func TestFormatValid(t *testing.T) {
	var (
		d        = DateFromTime(time.Now())
		expected = d.Time.Format("2006-01-02")
		str      = d.format()
	)
	if *str != expected {
		t.Errorf("Invalid format(). Actual: %s, expected: %s", *str, expected)

	}
}

func TestFormatInvalid(t *testing.T) {
	var (
		d   = Date{}
		str = d.format()
	)
	if str != nil {
		t.Errorf("Invalid format(). Actual: %s, expected: nil", *str)

	}
}

var testCaseParseString = []struct {
	value         string
	expectedError bool
}{
	// not time
	{
		value:         "invalid-format",
		expectedError: true,
	},
	// invalid format not rfc3339 w/a time
	{
		value:         "10-01-1996",
		expectedError: true,
	},
	// full time instead of date
	{
		value:         "1996-09-29T00:00:00",
		expectedError: true,
	},
	// only time instead of date
	{
		value:         "00:00:00",
		expectedError: true,
	},
	{
		value:         "",
		expectedError: false,
	},

	{
		value:         "1998-02-20",
		expectedError: false,
	},
}

func TestParseDateFromString(t *testing.T) {
	for _, testCase := range testCaseParseString {
		_, actualErr := parseDateFromString(testCase.value)
		if (actualErr == nil) == testCase.expectedError {
			t.Errorf(`Expecting error - '%v' but was opposite`, testCase.expectedError)
		}
	}
}

var testCaseString = []struct {
	date           Date
	expectedString string
}{
	{
		date:           Date{},
		expectedString: "",
	},
	{
		date:           DateFromTime(time.Date(1996, 8, 30, 10, 11, 40, 0, time.UTC)),
		expectedString: "1996-08-30",
	},
}

func TestString(t *testing.T) {
	for _, testCase := range testCaseString {
		// empty var d = Date{}
		var str = testCase.date.String()
		if str != testCase.expectedString {
			t.Errorf("Invalid String(). Actual: %s, expected: %s", str, testCase.expectedString)
		}
	}
}

func TestValuer(t *testing.T) {
	for _, testCase := range testCaseString {
		// empty var d = Date{}
		var _, err = testCase.date.Value()
		if err != nil {
			t.Errorf("Invalid String(). Err: %v", err)
		}
	}
}

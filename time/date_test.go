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
	testUnmarshal(t,
		func(bytes []byte, d *dt) error {
			return json.Unmarshal(bytes, d)
		}, func(u unmarshalType) []byte {
			return u.testJSON
		})
}

func TestUnmarshalXML(t *testing.T) {
	testUnmarshal(t,
		func(bytes []byte, d *dt) error {
			return xml.Unmarshal(bytes, d)
		}, func(u unmarshalType) []byte {
			return u.testXML
		})
}

func testUnmarshal(t *testing.T, unmarshalFunc func([]byte, *dt) error, valueFunc func(unmarshalType) []byte) {
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
	testMarshal(t, "json",
		func(d dt) ([]byte, error) {
			return json.Marshal(d)
		}, func(m marshalType) []byte {
			return m.expectedJSON
		})
}

func TestMarshalXML(t *testing.T) {
	testMarshal(t, "xml",
		func(d dt) ([]byte, error) {
			return xml.Marshal(d)
		}, func(m marshalType) []byte {
			return m.expectedXML
		})
}

func testMarshal(t *testing.T, tp string, marshalFunc func(dt) ([]byte, error), valueFunc func(marshalType) []byte) {
	for _, testCase := range marshalTestCases {
		actual, actualErr := marshalFunc(testCase.testData)
		if actualErr != nil {
			t.Errorf(`%s`, actualErr)
		}
		if string(actual) != string(valueFunc(testCase)) {
			t.Errorf("Actual %s does not equal expected. Actual: '%s', expected: '%s'", tp, actual, testCase.expectedXML)
		}
	}
}

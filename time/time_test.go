package time

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"
)

type tm struct {
	XMLName xml.Name `json:"-" xml:"test"`
	T       Time     `json:"t,omitempty" xml:"t,omitempty"`
}

var unmarshalTimeTestCases = []unmarshalType{
	{
		testJSON:               []byte(`{}`),
		testXML:                []byte(`<?xml version="1.0"?><test></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"t": ""}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t></t></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"t": "2022-05-10T12:35:10.132123Z"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t>2022-05-10T12:35:10.1341231Z</t></test>`),
		expectingValidateError: false,
	},
	{
		testJSON:               []byte(`{"t": 1234}`),
		testXML:                []byte(`test><t>@</t>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"t": "@"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t>@</t></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"t": "2006-01-02"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t>2006-01-02</t></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"t": "00:00:00"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t>00:00:00</t></test>`),
		expectingValidateError: true,
	},
	{
		testJSON:               []byte(`{"t": "text-text-text"}`),
		testXML:                []byte(`<?xml version="1.0"?><test><t>text-text-text</t></test>`),
		expectingValidateError: true,
	},
}

func TestTimeUnmarshalJSON(t *testing.T) {
	testTimeUnmarshal(t, json.Unmarshal, func(u unmarshalType) []byte {
		return u.testJSON
	})
}

func TestTimeUnmarshalXML(t *testing.T) {
	testTimeUnmarshal(t, xml.Unmarshal, func(u unmarshalType) []byte {
		return u.testXML
	})
}

func testTimeUnmarshal(t *testing.T, unmarshalFunc func([]byte, interface{}) error, valueFunc func(unmarshalType) []byte) {
	for _, testCase := range unmarshalTimeTestCases {
		var d tm
		actualErr := unmarshalFunc(valueFunc(testCase), &d)
		if (actualErr == nil) == testCase.expectingValidateError {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, valueFunc(testCase), testCase.expectingValidateError)
		}
	}
}

type marshalTimeType struct {
	testData     tm
	expectedJSON []byte
	expectedXML  []byte
}

var marshalTimeTestCases = []marshalTimeType{
	{
		testData:     tm{},
		expectedJSON: []byte(`{"t":null}`),
		expectedXML:  []byte(`<test></test>`),
	},
	{
		testData:     tm{T: NullTime(time.Date(2006, 1, 2, 0, 0, 0, 123456789, time.UTC))},
		expectedJSON: []byte(`{"t":"2006-01-02T00:00:00.123456789Z"}`),
		expectedXML:  []byte(`<test><t>2006-01-02T00:00:00.123456789Z</t></test>`),
	},
	{
		testData:     tm{T: NullTime(time.Date(1996, 8, 30, 10, 11, 40, 0, time.UTC))},
		expectedJSON: []byte(`{"t":"1996-08-30T10:11:40Z"}`),
		expectedXML:  []byte(`<test><t>1996-08-30T10:11:40Z</t></test>`),
	},
}

func TestTimeMarshalJSON(t *testing.T) {
	testTimeMarshal(t, "json", json.Marshal, func(m marshalTimeType) []byte {
		return m.expectedJSON
	})
}

func TestTimeMarshalXML(t *testing.T) {
	testTimeMarshal(t, "xml", xml.Marshal, func(m marshalTimeType) []byte {
		return m.expectedXML
	})
}

func testTimeMarshal(t *testing.T, tp string, marshalFunc func(interface{}) ([]byte, error), valueFunc func(timeType marshalTimeType) []byte) {
	for _, testCase := range marshalTimeTestCases {
		actual, actualErr := marshalFunc(testCase.testData)
		if actualErr != nil {
			t.Errorf(`%s`, actualErr)
		}
		if string(actual) != string(valueFunc(testCase)) {
			t.Errorf("Actual %s does not equal expected. Actual: '%s', expected: '%s'", tp, actual, valueFunc(testCase))
		}
	}
}

func TestTimeFormatValid(t *testing.T) {
	var (
		d        = NullTime(time.Now())
		expected = d.Time.Format(time.RFC3339Nano)
		str      = d.format()
	)
	if *str != expected {
		t.Errorf("Invalid format(). Actual: %s, expected: %s", *str, expected)

	}
}

func TestTimeFormatInvalid(t *testing.T) {
	var (
		d   = Time{}
		str = d.format()
	)
	if str != nil {
		t.Errorf("Invalid format(). Actual: %s, expected: nil", *str)

	}
}

var testCaseTimeParseString = []struct {
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
	// date instead of date time
	{
		value:         "1996-09-29",
		expectedError: true,
	},
	// only time instead of date time
	{
		value:         "00:00:00",
		expectedError: true,
	},
	{
		value:         "",
		expectedError: false,
	},
	{
		value:         "1998-02-20T12:21:40Z",
		expectedError: false,
	},
	{
		value:         "1998-02-20T12:21:40.4235324Z",
		expectedError: false,
	},
}

func TestParseTimeFromString(t *testing.T) {
	for _, testCase := range testCaseTimeParseString {
		var tm = &Time{}
		actualErr := tm.parseTimeFromString(testCase.value)
		if (actualErr == nil) == testCase.expectedError {
			t.Errorf(`Expecting error - '%v' but was opposite`, testCase.expectedError)
		}
	}
}

var testCaseTimeString = []struct {
	t              Time
	expectedString string
}{
	{
		t:              Time{},
		expectedString: "",
	},
	{
		t:              NullTime(time.Date(1996, 8, 30, 10, 11, 40, 123456789, time.UTC)),
		expectedString: "1996-08-30T10:11:40.123456789Z",
	},
}

func TestTimeString(t *testing.T) {
	for _, testCase := range testCaseTimeString {
		// empty var d = Date{}
		var str = testCase.t.String()
		if str != testCase.expectedString {
			t.Errorf("Invalid String(). Actual: %s, expected: %s", str, testCase.expectedString)
		}
	}
}

func TestTimeValuer(t *testing.T) {
	for _, testCase := range testCaseTimeString {
		// empty var d = Date{}
		var _, err = testCase.t.Value()
		if err != nil {
			t.Errorf("Invalid String(). Err: %v", err)
		}
	}
}

func TestParseNullTime_EmptyStr(t *testing.T) {
	var v, err = ParseNullTimeFromString("")
	if err != nil {
		t.Errorf("No error expected. Err: %v", err)
	}
	var nullValue Time
	if v != nullValue {
		t.Errorf("null value expected. Time: %v", err)
	}
}

func TestParseNullTime_ValidStr(t *testing.T) {
	var v, err = ParseNullTimeFromString("1996-08-30T10:11:40.123456789Z")
	if err != nil {
		t.Errorf("No error expected. Err: %v", err)
	}
	var validValue = NullTime(time.Date(1996, 8, 30, 10, 11, 40, 123456789, time.UTC))
	if v != validValue {
		t.Errorf("null value expected. Time: %v", err)
	}
}

func TestParseNullTime_NotValidStr(t *testing.T) {
	var v, err = ParseNullTimeFromString("00:00:00")
	if err == nil {
		t.Error("Error expected")
	}
	if err.Error() != "cannot parse time '00:00:00' invalid format '2006-01-02T15:04:05.999999999Z07:00'" {
		t.Errorf("%v expected: Got: %v", "cannot parse time '00:00:00' invalid format '2006-01-02T15:04:05.999999999Z07:00'", err)
	}
	var nullValue Time
	if v != nullValue {
		t.Errorf("null value expected. Time: %v", err)
	}
}

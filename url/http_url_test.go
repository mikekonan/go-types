package url

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testHttpUrlCases = []struct {
	url         HttpURL
	isExpectErr bool
}{
	{"", true},
	{"http://google.com", false},
	{"https://google.com#top", false},
	{"http://google.fist.com", false},
	{"http://google.second.ORG", false},
	{"ftp://server.en", true},
	{"mailto://mail1.3server.en", true},
	{"aaa://server.en", true},
	{"acap://1server.a.en", true},
	{"attachment://1server.a.en", true},
	{"aw://host:1234/aw-server.com", true},
	{"http://[2001:db8:a0b:12f0::1]/index.html", false},
	{"ftp://[1200:0000:AB00:1234:0000:2552:7777:1313]", true},
	{"https://user:pass@[::1]:9093/a/b/c/?a=v#abc", false},
	{"user:pass@[::1]:9093/a/b/c/?a=v#abc", true},
	{"https://127.0.0.1/a/b/c?a=v&c=11d", false},
}

func TestHttpURLValidate(t *testing.T) {
	for _, testCase := range testHttpUrlCases {
		actualErr := testCase.url.Validate()
		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.url, testCase.isExpectErr)
		}
	}
}

func TestHttpURLJsonUnmarshal(t *testing.T) {
	for _, testCase := range testHttpUrlCases {
		jsonStr := fmt.Sprintf(`{"url":"%s"}`, testCase.url)
		var emailStruct struct {
			Email HttpURL `json:"url"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &emailStruct)
		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.url, testCase.isExpectErr)
		}
	}
}

func TestHttpURLValue(t *testing.T) {
	for _, testCase := range testHttpUrlCases {
		actualValue, actualErr := testCase.url.Value()

		if testCase.url == "" {
			if actualValue != "" || actualErr != nil {
				t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.url, false)
			}

			continue
		}

		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`Value: '%s'. expecting error - '%v' but was opposite`, testCase.url, testCase.isExpectErr)
			if actualValue != "" {
				t.Errorf(`Value: '%s'. value have to be empty!`, testCase.url)
			}
		}
	}
}

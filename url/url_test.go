package url

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testCases = []struct {
	url         URL
	isExpectErr bool
}{
	{"", true},
	{"http://foo.bar#com", false},
	{"http://foobar.com", false},
	{"https://foobar.com", false},
	{"foobar.com", false},
	{"http://foobar.coffee/", false},
	{"http://foobar.中文网/", false},
	{"http://foobar.org/", false},
	{"http://foobar.ORG", false},
	{"http://foobar.org:8080/", false},
	{"ftp://foobar.ru/", false},
	{"ftp.foo.bar", false},
	{"http://user:pass@www.foobar.com/", false},
	{"http://user:pass@www.foobar.com/path/file", false},
	{"http://127.0.0.1/", false},
	{"http://duckduckgo.com/?q=%2F", false},
	{"http://localhost:3000/", false},
	{"http://foobar.com/?foo=bar#baz=qux", false},
	{"http://foobar.com?foo=bar", false},
	{"http://www.xn--froschgrn-x9a.net/", false},
	{"http://foobar.com/a-", false},
	{"http://foobar.پاکستان/", false},
	{"http://foobar.c_o_m", true},
	{"http://_foobar.com", true},
	{"http://foo_bar.com", false},
	{"http://user:pass@foo_bar_bar.bar_foo.com", false},
	{"", true},
	{"xyz://foobar.com", true},
	{".com", true},
	{"rtmp://foobar.com", true},
	{"http://localhost:3000/", false},
	{"http://foobar.com#baz=qux", false},
	{"http://www.foobar.com/~foobar", false},
	{"http://www.-foobar.com/", true},
	{"http://www.foo---bar.com/", true},
	{"http://r6---snnvoxuioq6.googlevideo.com", false},
	{"mailto:someone@example.com", false},
	{"irc://irc.server.org/channel", true},
	{"irc://#channel@network", false},
	{"/abs/test/dir", true},
	{"./rel/test/dir", true},
	{"http://foo^bar.org", true},
	{"http://foo&*bar.org", true},
	{"http://foo&bar.org", true},
	{"http://foo bar.org", true},
	{"http://foo.bar.org", false},
	{"http://www.foo.bar.org", false},
	{"http://www.foo.co.uk", false},
	{"foo", true},
	{"http://.foo.com", true},
	{"http://,foo.com", true},
	{",foo.com", true},
	{"http://myservice.:9093/", false},
	{"https://pbs.twimg.com/profile_images/560826135676588032/j8fWrmYY_normal.jpeg", false},
	{"http://prometheus-alertmanager.service.q:9093", false},
	{"aio1_alertmanager_container-63376c45:9093", false},
	{"https://www.logn-123-123.url.with.sigle.letter.d:12345/url/path/foo?bar=zzz#user", false},
	{"http://me.example.com", false},
	{"http://www.me.example.com", false},
	{"https://farm6.static.flickr.com", false},
	{"https://zh.wikipedia.org/wiki/Wikipedia:%E9%A6%96%E9%A1%B5", false},
	{"google", true},
	{"http://hyphenated-host-name.example.co.in", false},
	{"http://cant-end-with-hyphen-.example.com", true},
	{"http://-cant-start-with-hyphen.example.com", true},
	{"http://www.domain-can-have-dashes.com", false},
	{"http://m.abcd.com/test.html", false},
	{"http://m.abcd.com/a/b/c/d/test.html?args=a&b=c", false},
	{"http://[::1]:9093", false},
	{"http://[::1]:909388", true},
	{"1200::AB00:1234::2552:7777:1313", true},
	{"http://[2001:db8:a0b:12f0::1]/index.html", false},
	{"http://[1200:0000:AB00:1234:0000:2552:7777:1313]", false},
	{"http://user:pass@[::1]:9093/a/b/c/?a=v#abc", false},
	{"https://127.0.0.1/a/b/c?a=v&c=11d", false},
	{"https://foo_bar.example.com", false},
	{"http://foo_bar.example.com", false},
	{"http://foo_bar_fizz_buzz.example.com", false},
	{"http://_cant_start_with_underescore", true},
	{"http://cant_end_with_underescore_", true},
	{"foo_bar.example.com", false},
	{"foo_bar_fizz_buzz.example.com", false},
	{"http://hello_world.example.com", false},
	{"foo_bar-fizz-buzz:1313", false},
	{"foo_bar-fizz-buzz:13:13", true},
	{"foo_bar-fizz-buzz://1313", true},
	{"https://stage-payflow.tonder.io/oxxo-pay?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc2Nzk1OTIsImlhdCI6MTc1NzY3Nzc5MiwiaXNzIjoiY2hlY2tvdXQtcm91dGVyIiwic2NvcGUiOiJjaGVja291dCIsImNoZWNrb3V0X2lkIjoiODZkYmE3NWEtYjFkNy00NWNmLWJhMjEtZTQzNmNmZDM3NjU2IiwicGF5bWVudF9pZCI6NDU3NjIsInRyYW5zYWN0aW9uX2lkIjoiODdkY2Y0OGUtMzkyNC00YjM5LThkNTctMDY4OTY1Y2FjZjNiIiwiYnVzaW5lc3NfaWQiOiIxMzYiLCJyZWRpcmVjdF91cmwiOiJodHRwczovL3VzcnYtYXBtcy1zdGFnZS1veHhvcGF5LXRpY2tldHMuczMudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vb3h4by1wYXltZW50LTk4MDAyMDE5Mjk3NTEzLnBkZj9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1Db250ZW50LVNoYTI1Nj1VTlNJR05FRC1QQVlMT0FEJlgtQW16LUNyZWRlbnRpYWw9QVNJQTRYVFdZSDVYVDRIS1hNT0olMkYyMDI1MDkxMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNTA5MTJUMTE0OTUyWiZYLUFtei1FeHBpcmVzPTg2NDAwJlgtQW16LVNlY3VyaXR5LVRva2VuPUlRb0piM0pwWjJsdVgyVmpFTFQlMkYlMkYlMkYlMkYlMkYlMkYlMkYlMkYlMkYlMkZ3RWFDWFZ6TFdWaGMzUXRNU0pITUVVQ0lRQ1gwbTVmVFBIZThEOFREN3NubHZyOE5ueUhQJTJCcXJwNUd2NCUyQiUyRmppUmJoQ2dJZ0JvRm9rY1U4cEtHNkp5MWsyWjVOdlNyUGkyMmhWY1RIdVFlT0xkSjclMkI4SXFqd01JTFJBQ0dndzROelV6TkRnMk16azFPVGtpRERkNlZncE12ajBVRlZSSHRDcnNBcWY3YXFFcnJ4ZDB4ZiUyRkhTTWlXVmFjd3BlaXRmOUhDdEtBemNxTWNRMkZIa1ZEd2VFVzQlMkZiQnolMkJoWmpvcCUyQlZkQm4wcDBHY3puUDFDMnBPOTg2RTlMYnc3QjBvYlpnSWtTYTFBemZxeVNCZHkzU1dFNm1RYzB3b0plZ0wyUGhEVGMlMkYlMkZDJTJGUmFSQmtUT2dnRHBUNFZtOG5ZWWJBMVNIT1NzNkdqRDdSVkZ5UkVVNWNiSGNkV1J4QzZ1MWFmOG5xN1VvdWp0VmF0OCUyQkVtcGdVeDFxWEtpQ21idGRJN2xJN1JSQm5xMlI5U3BqNnd3YUdoYTJQTld5RWV1b1A1eWJmJTJGQjQzVHpkJTJGV1c3ZHdnSHJ6ZVZnNUNvN3UzTnhrQ0Z5UjR2cllKb1hYMGc3OXJKeEcyTGd6S3FmemxsbXZyN0dsT1czS21MQmxmenFnY0dNa0tvcFNzVjMlMkZUcFNLVGp1SkhxZTRnYzc0UkRnT1VFcUo4aXM2WlA5UnNFRnNoVWszcDB1ZkFIbVhTSXM3N3hlNSUyRmVqWkV4U0lHTkR0MG94OWUzckdaZkpXeVRJU3NYRGFueFJ1V1ZkSFphdkFsbGtnbWRHc0cyWTJLdFN2djY3bHhVZnNuWXE2S3RxOGtrbUdWTlFJQ05rdzNKR1F4Z1k2blFGZzE3eExLMjE1Vmp6WVdtS2V6bFFzUUx1bkFMSk1PZWdVTHo0VDUxZUJhJTJGUmxaS041YkFzYTRqblJORDQ5R0FyU2Z2MkVuOWdoTHpxNGFIaHZkZE5pVDNIN2I1djVra1RJVTJFMnhqams1ZXJyRlFzNjBPJTJGTUo3VE1ReG9TR3hKOG1FbGIweDhoUlZiUGVrTGl5NXA0Zm1HUE96VnYzT2J1N1k1UUpHNk9BS1NQS1c5dXlqSnlmYXljb1B6WmxsbUpnNXRmcDk3TTdxOHdWS2cxJlgtQW16LVNpZ25hdHVyZT0yODExMmMxNTQ2NWEyMjBiYmMwODExYjE4MmZiMTc5Y2I4MThhNjlmZDM2ZDNjMWYzNWMxZjAwODU5ODQ4ZTMyJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZ4LWFtei1jaGVja3N1bS1tb2RlPUVOQUJMRUQmeC1pZD1HZXRPYmplY3QiLCJyZXR1cm5fdXJsIjoiaHR0cHM6Ly9nb29nbGUuY29tP3RuZHJfcGF5bWVudF9pZD00NTc2MiIsInByb3ZpZGVyIjoib3h4b3BheSIsInBheW1lbnRfbWV0aG9kIjoib3h4b3BheSIsImlzX3JvdXRlX2ZpbmlzaGVkIjp0cnVlLCJleHRyYV9kYXRhIjp7ImJhcmNvZGUiOiI5ODAwMjAxOTI5NzUxMyIsImV4cGlyZWQiOiIxNzYwMjI3MjAwIiwiYW1vdW50Ijo1MDAuMH19.mAVsqYx5nD974pCk-RWnYoUCNnqbXuejbK5E2tsuS6g", false},
}

func TestURLValidate(t *testing.T) {
	for _, testCase := range testCases {
		actualErr := testCase.url.Validate()
		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`Validate: '%s'. expecting error - '%v' but was opposite`, testCase.url, testCase.isExpectErr)
		}
	}
}

func TestURLJsonUnmarshal(t *testing.T) {
	for _, testCase := range testCases {
		jsonStr := fmt.Sprintf(`{"url":"%s"}`, testCase.url)
		var emailStruct struct {
			Email URL `json:"url"`
		}

		actualErr := json.Unmarshal([]byte(jsonStr), &emailStruct)
		if (actualErr == nil) == testCase.isExpectErr {
			t.Errorf(`JsonUnmarshal: '%s'. expecting error - '%v' but was opposite`, testCase.url, testCase.isExpectErr)
		}
	}
}

func TestURLValue(t *testing.T) {
	for _, testCase := range testCases {
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

/*
Package url provides simple URL type.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package url

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	// Most major browsers have practical URL length limits around 2,048 characters. Here are some specifics:
	// - Internet Explorer: 2,083 characters (2,048 for path/query, plus 35 for the protocol/host)
	// - Edge: ~2,048 characters
	// - Chrome: ~32,768 characters (but compatibility issues may arise above 2,048)
	// - Firefox: ~65,536 characters (but practical compatibility is around 2,048)
	// - Safari: ~80,000 characters (but practical compatibility is around 2,048)
	// - Opera: ~190,000 characters (but practical compatibility is around 2,048)
	maxURLRuneCount = 3000
	minURLRuneCount = 3

	PatternIP           string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	PatternURLSchema    string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	PatternURLUsername  string = `(\S+(:\S*)?@)`
	PatternURLPort      string = `(:(\d{1,5}))`
	PatternURLPath      string = `((\/|\?|#)[^\s]*)`
	PatternURLIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	PatternURLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	PatternURL                 = `^` + PatternURLSchema + `?` + PatternURLUsername + `?` + `((` + PatternURLIP + `|(\[` + PatternIP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + PatternURLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + PatternURLPort + `?` + PatternURLPath + `?$`
)

var rxURL = regexp.MustCompile(PatternURL)

// URL represents a URL type
type URL string

// Value implementation of driver.Valuer
func (url URL) Value() (value driver.Value, err error) {
	if url == "" {
		return "", nil
	}

	if err = url.Validate(); err != nil {
		return nil, err
	}

	return url.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (url URL) Validate() error {
	if !isURL(url.String()) {
		return fmt.Errorf("'%s' is not a valid url", url)
	}

	return nil
}

// IsURL checks if the string is an URL.
// Copy function https://github.com/asaskevich/govalidator/blob/e11347878e2323a0777b40d33eeffd37a185e31b/validator.go#L104
// As we have cases when maxURLRuneCount is exceeded
func isURL(str string) bool {
	if str == "" || utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount || strings.HasPrefix(str, ".") {
		return false
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return rxURL.MatchString(str)
}

// UnmarshalJSON unmarshall implementation for Email
func (url *URL) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	value := URL(str)
	if err := value.Validate(); err != nil {
		return err
	}

	*url = value

	return nil
}

// String implementation of Stringer interface
func (url URL) String() string {
	return string(url)
}

// HttpURL represents a URL type with http/https schema
type HttpURL string

// Value implementation of driver.Valuer
func (httpUrl HttpURL) Value() (value driver.Value, err error) {
	if httpUrl == "" {
		return "", nil
	}

	if err = httpUrl.Validate(); err != nil {
		return nil, err
	}

	return httpUrl.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (httpUrl HttpURL) Validate() (err error) {
	if len(httpUrl) == 0 {
		return fmt.Errorf("'%s' is not a valid http url", httpUrl)
	}

	if err = validateSchema(httpUrl.String()); err != nil {
		return err
	}

	return URL(httpUrl).Validate()
}

// UnmarshalJSON unmarshall implementation for Email
func (httpUrl *HttpURL) UnmarshalJSON(data []byte) error {
	var url string
	if err := json.Unmarshal(data, &url); err != nil {
		return err
	}

	value := HttpURL(url)
	if err := value.Validate(); err != nil {
		return err
	}

	*httpUrl = value

	return nil
}

// String implementation of Stringer interface
func (httpUrl HttpURL) String() string {
	return string(httpUrl)
}

// NullHttpURL represents a URL type with http/https schema
// Can be empty
type NullHttpURL string

// Value implementation of driver.Valuer
func (nullHttpURL NullHttpURL) Value() (value driver.Value, err error) {
	if nullHttpURL == "" {
		return "", nil
	}

	if err = nullHttpURL.Validate(); err != nil {
		return nil, err
	}

	return nullHttpURL.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (nullHttpURL NullHttpURL) Validate() (err error) {
	if len(nullHttpURL) == 0 {
		return nil
	}

	if err = validateSchema(nullHttpURL.String()); err != nil {
		return err
	}

	return URL(nullHttpURL).Validate()
}

// UnmarshalJSON unmarshall implementation for Email
func (nullHttpURL *NullHttpURL) UnmarshalJSON(data []byte) error {
	var url string
	if err := json.Unmarshal(data, &url); err != nil {
		return err
	}

	value := NullHttpURL(url)
	if err := value.Validate(); err != nil {
		return err
	}

	*nullHttpURL = value

	return nil
}

// String implementation of Stringer interface
func (nullHttpURL NullHttpURL) String() string {
	return string(nullHttpURL)
}

func validateSchema(url string) error {
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		return fmt.Errorf("'%s' is not a valid http schema", url)
	}

	return nil
}

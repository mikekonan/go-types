/*
Package url provides simple URL type.

Author Mikalai Konan(mikalai.konan@icloud.com).
*/
package url

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

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
	if !govalidator.IsURL(url.String()) {
		return fmt.Errorf("'%s' is not a valid url", url)
	}

	return nil
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

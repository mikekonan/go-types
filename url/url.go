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

//Value implementation of driver.Valuer
func (url URL) Value() (value driver.Value, err error) {
	if url == "" {
		return "", nil
	}

	if err = url.Validate(); err != nil {
		return nil, err
	}

	return url.String(), nil
}

//Validate implementation of ozzo-validation Validate interface
func (url URL) Validate() error {
	if !govalidator.IsURL(url.String()) {
		return fmt.Errorf("'%s' is not a valid url", url)
	}

	return nil
}

//UnmarshalJSON unmarshall implementation for Email
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

//String implementation of Stringer interface
func (url URL) String() string {
	return string(url)
}

// HttpURL represents a URL type with http/https schema
type HttpURL string

//Value implementation of driver.Valuer
func (httpUrl HttpURL) Value() (value driver.Value, err error) {
	if httpUrl == "" {
		return "", nil
	}

	if err = httpUrl.Validate(); err != nil {
		return nil, err
	}

	return httpUrl.String(), nil
}

//Validate implementation of ozzo-validation Validate interface
//Url can be empty (use required validation, if cannot be empty)
func (httpUrl HttpURL) Validate() error {
	if len(httpUrl) == 0 {
		return nil
	}

	if !strings.HasPrefix(httpUrl.String(), "http:") && !strings.HasPrefix(httpUrl.String(), "https:") {
		return fmt.Errorf("'%s' is not a valid http schema", httpUrl)
	}

	return URL(httpUrl).Validate()
}

//UnmarshalJSON unmarshall implementation for Email
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

//String implementation of Stringer interface
func (httpUrl HttpURL) String() string {
	return string(httpUrl)
}

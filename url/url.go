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
		return fmt.Errorf("'%s' is not valid url", url)
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

// URL represents a URL type with schema
type HttpURL string

//Value implementation of driver.Valuer
func (url HttpURL) Value() (value driver.Value, err error) {
	if url == "" {
		return "", nil
	}

	if err = url.Validate(); err != nil {
		return nil, err
	}

	return url.String(), nil
}

//Validate implementation of ozzo-validation Validate interface
func (url HttpURL) Validate() error {
	if !strings.HasPrefix(url.String(), "http:") || !strings.HasPrefix(url.String(), "https:") {
		return fmt.Errorf("'%s' is not valid url schema", url)
	}

	if !govalidator.IsRequestURL(url.String()) {
		return fmt.Errorf("'%s' is not valid url", url)
	}

	return nil
}

//UnmarshalJSON unmarshall implementation for Email
func (url *HttpURL) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	value := HttpURL(str)
	if err := value.Validate(); err != nil {
		return err
	}

	*url = value

	return nil
}

//String implementation of Stringer interface
func (url HttpURL) String() string {
	return string(url)
}

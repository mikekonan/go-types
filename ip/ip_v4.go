package ip

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/netip"
)

type IPv4 string

// Value implementation of driver.Valuer
func (ip IPv4) Value() (value driver.Value, err error) {
	if err = ip.Validate(); err != nil {
		return nil, err
	}

	return ip.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (ip IPv4) Validate() error {
	_, err := V4FromString(ip.String())
	if err != nil {
		return fmt.Errorf("'%s' is not an valid IPv4 address, err: %w", ip, err)
	}

	return nil
}

// UnmarshalJSON unmarshall implementation for IP
func (ip *IPv4) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	value := IPv4(str)
	if err := value.Validate(); err != nil {
		return err
	}

	*ip = value

	return nil
}

// String implementation of Stringer interface
func (ip IPv4) String() string {
	return string(ip)
}

func V4FromString(value string) (IPv4, error) {
	if value == "" {
		return "", ErrEmptyValue
	}

	var addr, err = netip.ParseAddr(value)
	if err != nil {
		return "", err
	}

	if !addr.Is4() {
		return "", fmt.Errorf("not an IPv4 address")
	}

	return IPv4(addr.String()), nil
}

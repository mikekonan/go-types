package ip

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/netip"
)

type IPv6 string

// Value implementation of driver.Valuer
func (ip IPv6) Value() (value driver.Value, err error) {
	if err = ip.Validate(); err != nil {
		return nil, err
	}

	return ip.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (ip IPv6) Validate() error {
	_, err := V6FromString(ip.String())
	if err != nil {
		return validationError(ip, "IPv6", err)
	}

	return nil
}

// UnmarshalJSON unmarshall implementation for IP
func (ip *IPv6) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str != "" {
		value, err := V6FromString(str)
		if err != nil {
			return err
		}

		*ip = value
	}

	return nil
}

// String implementation of Stringer interface
func (ip IPv6) String() string {
	return string(ip)
}

func V6FromString(value string) (IPv6, error) {
	if value == "" {
		return "", fmt.Errorf("empty value")
	}

	var addr, err = netip.ParseAddr(value)
	if err != nil {
		return "", err
	}

	if !addr.Is6() {
		return "", fmt.Errorf("not an IPv6 address")
	}

	return IPv6(addr.String()), nil
}

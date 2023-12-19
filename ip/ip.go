package ip

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type IP struct {
	v4   IPv4
	v6   IPv6
	raw  string
	isV6 bool
}

// Value implementation of driver.Valuer
func (ip IP) Value() (value driver.Value, err error) {
	if err = ip.Validate(); err != nil {
		return nil, err
	}

	return ip.String(), nil
}

// Validate implementation of ozzo-validation Validate interface
func (ip IP) Validate() error {
	var fromStringIP, err = FromString(ip.String())
	if err != nil {
		return fmt.Errorf("'%s' is not an valid IP address, err: %w", ip, err)
	}

	if fromStringIP.raw == ip.raw &&
		fromStringIP.isV6 == ip.isV6 &&
		fromStringIP.v4 == ip.v4 &&
		fromStringIP.v6 == ip.v6 {
		return nil
	}

	return errors.New(`IP inconsistency error`)
}

// UnmarshalJSON unmarshall implementation for IP
func (ip *IP) UnmarshalJSON(data []byte) error {
	var (
		str string
		err error
	)
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	var value IP
	if value, err = FromString(str); err != nil && value.Validate() != nil {
		return err
	}

	*ip = value

	return nil
}

// MarshalJSON marshall implementation for IP
func (ip IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.raw)
}

// String implementation of Stringer interface
func (ip IP) String() string {
	return ip.raw
}

func (ip IP) IPv4() IPv4 {
	return ip.v4
}

func (ip IP) IPv6() IPv6 {
	return ip.v6
}

func (ip IP) IsIPv6() bool {
	return ip.isV6
}

func FromString(value string) (IP, error) {
	if value == "" {
		return IP{}, fmt.Errorf("empty value")
	}

	var v4Addr, err = V4FromString(value)
	if err == nil {
		return IP{
			v4:   v4Addr,
			isV6: false,
			raw:  value,
		}, nil
	}

	var v6Addr IPv6
	if v6Addr, err = V6FromString(value); err == nil {
		return IP{
			v6:   v6Addr,
			isV6: true,
			raw:  value,
		}, nil
	}

	return IP{}, errors.New("value isn't fit neither v4 nor v6 IP")
}

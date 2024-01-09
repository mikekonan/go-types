package ip

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestIP_Value(t *testing.T) {
	tests := []struct {
		name      string
		ip        IP
		wantValue driver.Value
		wantErr   bool
	}{
		{
			name: "IPv4",
			ip: IP{
				isV6: false,
				v4:   IPv4("1.1.1.1"),
				raw:  "1.1.1.1",
			},
			wantValue: "1.1.1.1",
			wantErr:   false,
		},
		{
			name: "invalid-IPv4",
			ip: IP{
				isV6: false,
				v4:   IPv4("2001:0db8:85a3:0000:0000:8a2e:0370:73341"),
				raw:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "invalid-both-v6false",
			ip: IP{
				v4:  IPv4("1111"),
				v6:  IPv6("1111"),
				raw: "1111",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "invalid-both-v6true",
			ip: IP{
				isV6: true,
				v4:   IPv4("1111"),
				v6:   IPv6("1111"),
				raw:  "1111",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "invalid-v4",
			ip: IP{
				v4:  IPv4("1111"),
				raw: "1111",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "invalid-v6",
			ip: IP{
				isV6: true,
				v6:   IPv6("1111"),
				raw:  "1111",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "IPv6",
			ip: IP{
				isV6: true,
				v6:   IPv6("2001:db8:85a3::8a2e:370:7334"),
				raw:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantValue: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr:   false,
		},
		{
			name: "invalid-IPv6",
			ip: IP{
				isV6: true,
				v6:   IPv6("1.1.1.1"),
				raw:  "1.1.1.1",
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "empty",
			ip: IP{
				raw: "",
			},
			wantValue: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := tt.ip.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Value() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestIP_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name              string
		ip                string
		wantIP            IP
		wantErr           error
		wantValidationErr error
	}{
		{
			name:   "IPv4",
			ip:     "1.1.1.1",
			wantIP: IP{v4: "1.1.1.1", raw: "1.1.1.1"},
		},
		{
			name:    "invalid",
			ip:      "1111",
			wantErr: ErrInvalidIPFormat,
		},
		{
			name:   "IPv6",
			ip:     "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantIP: IP{v6: "2001:db8:85a3::8a2e:370:7334", raw: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", isV6: true},
		},
		{
			name:              "empty",
			ip:                "",
			wantValidationErr: fmt.Errorf("'' is not a valid IP address: empty value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ip IP
			err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, tt.ip)), &ip)

			if (err == nil) != (tt.wantErr == nil) || (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("%s: got error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}

			if tt.wantErr != nil {
				return
			}

			if !reflect.DeepEqual(ip, tt.wantIP) {
				t.Errorf("%s: got IP = %#v, want %#v", tt.name, ip, tt.wantIP)
				return
			}

			err = ip.Validate()
			if (err == nil) != (tt.wantValidationErr == nil) || (err != nil && tt.wantValidationErr != nil && err.Error() != tt.wantValidationErr.Error()) {
				t.Errorf("%s: unexpected validation error = %v, want %v", tt.name, err, tt.wantValidationErr)
				return
			}
		})
	}
}

func TestIP_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ip      IP
		wantErr bool
	}{
		{
			name: "IPv4",
			ip: IP{
				isV6: false,
				v4:   IPv4("1.1.1.1"),
				raw:  "1.1.1.1",
			},
			wantErr: false,
		},
		{
			name: "invalid-IPv4",
			ip: IP{
				isV6: false,
				v4:   IPv4("2001:0db8:85a3:0000:0000:8a2e:0370:73341"),
				raw:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantErr: true,
		},
		{
			name: "invalid-both-v6false",
			ip: IP{
				v4:  IPv4("1111"),
				v6:  IPv6("1111"),
				raw: "1111",
			},
			wantErr: true,
		},
		{
			name: "invalid-both-v6true",
			ip: IP{
				isV6: true,
				v4:   IPv4("1111"),
				v6:   IPv6("1111"),
				raw:  "1111",
			},
			wantErr: true,
		},
		{
			name: "invalid-v4",
			ip: IP{
				v4:  IPv4("1111"),
				raw: "1111",
			},
			wantErr: true,
		},
		{
			name: "invalid-v6",
			ip: IP{
				isV6: true,
				v6:   IPv6("1111"),
				raw:  "1111",
			},
			wantErr: true,
		},
		{
			name: "IPv6",
			ip: IP{
				isV6: true,
				v6:   IPv6("2001:db8:85a3::8a2e:370:7334"),
				raw:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantErr: false,
		},
		{
			name: "invalid-IPv6",
			ip: IP{
				isV6: true,
				v6:   IPv6("1.1.1.1"),
				raw:  "1.1.1.1",
			},
			wantErr: true,
		},
		{
			name: "empty",
			ip: IP{
				raw: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ip.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseIPFromString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue IP
		wantErr   bool
	}{
		{
			name:  "IPv4",
			value: "1.1.1.1",
			wantValue: IP{
				isV6: false,
				v4:   IPv4("1.1.1.1"),
				raw:  "1.1.1.1",
			},
			wantErr: false,
		},
		{
			name:      "invalid",
			value:     "1111",
			wantValue: IP{},
			wantErr:   true,
		},
		{
			name:  "IPv6",
			value: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantValue: IP{
				isV6: true,
				v6:   IPv6("2001:db8:85a3::8a2e:370:7334"),
				raw:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantErr: false,
		},
		{
			name:      "empty",
			value:     "",
			wantValue: IP{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIPFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantValue {
				t.Errorf("ParseIPFromString() got = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

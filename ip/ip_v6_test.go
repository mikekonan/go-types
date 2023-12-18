package ip

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

func TestIPv6_Value(t *testing.T) {
	tests := []struct {
		name      string
		ip        IPv6
		wantValue driver.Value
		wantErr   bool
	}{
		{
			name:      "IPv4",
			ip:        "1.1.1.1",
			wantValue: nil,
			wantErr:   true,
		},
		{
			name:      "invalid",
			ip:        "1111",
			wantValue: nil,
			wantErr:   true,
		},
		{
			name:      "IPv6",
			ip:        "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantValue: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr:   false,
		},
		{
			name:      "empty",
			ip:        "",
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

func TestIPv6_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		wantErr bool
	}{
		{
			name:    "IPv4",
			ip:      "1.1.1.1",
			wantErr: true,
		},
		{
			name:    "invalid",
			ip:      "1111",
			wantErr: true,
		},
		{
			name:    "IPv6",
			ip:      "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: false,
		},
		{
			name:    "empty",
			ip:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ipv6 IPv6
			err := json.Unmarshal([]byte(tt.ip), &ipv6)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPv6_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ip      IPv6
		wantErr bool
	}{
		{
			name:    "IPv4",
			ip:      "1.1.1.1",
			wantErr: true,
		},
		{
			name:    "invalid",
			ip:      "1111",
			wantErr: true,
		},
		{
			name:    "IPv6",
			ip:      "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: false,
		},
		{
			name:    "empty",
			ip:      "",
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

func TestParseIPv6FromString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue IPv6
		wantErr   bool
	}{
		{
			name:      "IPv4",
			value:     "1.1.1.1",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "invalid",
			value:     "1111",
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "IPv6",
			value:     "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantValue: "2001:db8:85a3::8a2e:370:7334",
			wantErr:   false,
		},
		{
			name:      "empty",
			value:     "",
			wantValue: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := V6FromString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIPv6FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantValue {
				t.Errorf("ParseIPv6FromString() got = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

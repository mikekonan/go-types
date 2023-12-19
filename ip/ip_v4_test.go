package ip

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestIPv4_Value(t *testing.T) {
	tests := []struct {
		name      string
		ip        IPv4
		wantValue driver.Value
		wantErr   bool
	}{
		{
			name:      "success",
			ip:        "1.1.1.1",
			wantValue: "1.1.1.1",
			wantErr:   false,
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
			wantValue: nil,
			wantErr:   true,
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

func TestIPv4_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		ip      IPv4
		wantErr bool
	}{
		{
			name:    "success",
			ip:      "1.1.1.1",
			wantErr: false,
		},
		{
			name:    "invalid",
			ip:      "1111",
			wantErr: true,
		},
		{
			name:    "IPv6",
			ip:      "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: true,
		},
		{
			name:    "empty",
			ip:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonIPv4 := []byte(fmt.Sprintf(`{"IPv4": "%s"}`, tt.ip))
			var ipv4 struct {
				IPv4 IPv4 `json:"IPv4"`
			}

			err := json.Unmarshal(jsonIPv4, &ipv4)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPv4_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ip      IPv4
		wantErr bool
	}{
		{
			name:    "success",
			ip:      "1.1.1.1",
			wantErr: false,
		},
		{
			name:    "invalid",
			ip:      "1111",
			wantErr: true,
		},
		{
			name:    "IPv6",
			ip:      "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: true,
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

func TestParseIPv4FromString(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantValue IPv4
		wantErr   bool
	}{
		{
			name:      "success",
			value:     "1.1.1.1",
			wantValue: "1.1.1.1",
			wantErr:   false,
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
			wantValue: "",
			wantErr:   true,
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
			got, err := V4FromString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIPv4FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantValue {
				t.Errorf("ParseIPv4FromString() got = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

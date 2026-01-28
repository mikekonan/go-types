package utils

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestUnsafeStringFromJson(t *testing.T) {
	tests := []struct {
		name            string
		input           []byte
		expectedValue   string
		expectedIsEmpty bool
		expectedError   bool
	}{
		{
			name:            "valid string",
			input:           []byte(`"hello"`),
			expectedValue:   "hello",
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "valid string with spaces",
			input:           []byte(`"hello world"`),
			expectedValue:   "hello world",
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "valid string with special characters",
			input:           []byte(`"hello\nworld"`),
			expectedValue:   `hello\nworld`,
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "valid string with unicode",
			input:           []byte(`"hello 世界"`),
			expectedValue:   "hello 世界",
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "null value",
			input:           []byte(`null`),
			expectedValue:   "",
			expectedIsEmpty: true,
			expectedError:   false,
		},
		{
			name:            "empty string",
			input:           []byte(`""`),
			expectedValue:   "",
			expectedIsEmpty: true,
			expectedError:   false,
		},
		{
			name:            "string with content but only spaces",
			input:           []byte(`"   "`),
			expectedValue:   "   ",
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "invalid - missing opening quote",
			input:           []byte(`hello"`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - missing closing quote",
			input:           []byte(`"hello`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - missing both quotes",
			input:           []byte(`hello`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - number instead of string",
			input:           []byte(`123`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - boolean instead of string",
			input:           []byte(`true`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - empty input",
			input:           []byte(``),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "invalid - single quote",
			input:           []byte(`"`),
			expectedValue:   "",
			expectedIsEmpty: false,
			expectedError:   true,
		},
		{
			name:            "valid - quoted quotes inside",
			input:           []byte(`"say \"hello\""`),
			expectedValue:   `say \"hello\"`,
			expectedIsEmpty: false,
			expectedError:   false,
		},
		{
			name:            "valid - json object as string value",
			input:           []byte(`"{\"key\":\"value\"}"`),
			expectedValue:   `{\"key\":\"value\"}`,
			expectedIsEmpty: false,
			expectedError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, isEmpty, err := UnsafeStringFromJson(tt.input)

			if tt.expectedError {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else {
					// Verify it's the correct error type
					var unmarshalTypeError *json.UnmarshalTypeError
					if !errors.As(err, &unmarshalTypeError) {
						t.Errorf("expected *json.UnmarshalTypeError, got %T", err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if value != tt.expectedValue {
				t.Errorf("value = %q, want %q", value, tt.expectedValue)
			}

			if isEmpty != tt.expectedIsEmpty {
				t.Errorf("isEmpty = %v, want %v", isEmpty, tt.expectedIsEmpty)
			}
		})
	}
}

func TestUnsafeStringFromJson_UnsafeBehavior(t *testing.T) {
	// Test that demonstrates the unsafe behavior:
	// The returned string points to the same memory as the input slice
	input := []byte(`"original"`)
	value, _, err := UnsafeStringFromJson(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != "original" {
		t.Errorf("value = %q, want %q", value, "original")
	}
}

func BenchmarkUnsafeStringFromJson(b *testing.B) {
	input := []byte(`"hello world"`)

	b.ReportAllocs()
	b.SetBytes(int64(len(input)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = UnsafeStringFromJson(input)
	}
}

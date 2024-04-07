package validate_test

import (
	"cbam_api/validate"
	"fmt"
	"testing"
)

func TestValidateDecimals(t *testing.T) {
	type testCase struct {
		value     interface{}
		typeName  string
		expectErr bool
	}

	tests := []testCase{

		{value: 12.345, typeName: "DECIMAL900", expectErr: false},
		{value: 123.45, typeName: "DECIMAL5", expectErr: false},
		{value: 12345.67, typeName: "DECIMAL7", expectErr: false},

		{value: 12345.67, typeName: "DECIMAL0.5", expectErr: true},
		{value: 12345.67, typeName: "DECIMAL4", expectErr: true},
		{value: 12345.67, typeName: "DECIMAL-1", expectErr: true},
		{value: 12345.6700, typeName: "DECIMAL3", expectErr: true},
		{value: "not a float", typeName: "DECIMAL1", expectErr: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ValidateDecimals(%v, %s)", test.value, test.typeName), func(t *testing.T) {
			result, err := validate.ValidateDecimals(test.value, test.typeName)
			if test.expectErr && err == nil {
				t.Errorf("expected error but got none")
			}
			if !test.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if test.expectErr && err != nil && result {
				t.Errorf("validation passed but an error was expected: %v", err)
			}
		})
	}
}

func TestFormatFloatToString(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{input: 12.34, expected: "12.34"},
		{input: 123.456, expected: "123.456"},
		{input: 12345.6789, expected: "12345.6789"},
		{input: -12.34, expected: "-12.34"},
		{input: -123.456, expected: "-123.456"},
		{input: -12345.6789, expected: "-12345.6789"},
	}

	for _, test := range tests {
		t.Run("FormatFloatToString", func(t *testing.T) {
			result := validate.FormatFloatToString(test.input)
			if result != test.expected {
				t.Errorf("Expected %s but got %s for input %f", test.expected, result, test.input)
			}
		})
	}
}

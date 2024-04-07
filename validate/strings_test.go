package validate_test

import (
	"cbam_api/validate"
	"fmt"
	"testing"
)

func TestValidateStrings(t *testing.T) {
	type testCase struct {
		value     interface{}
		typeName  string
		expectErr bool
	}

	tests := []testCase{

		{value: "TESTSMALL", typeName: "STRING900", expectErr: false},
		{value: "TESTS", typeName: "STRING5", expectErr: false},
		{value: "TESTSMA", typeName: "STRING7", expectErr: false},

		{value: "TESTSMA", typeName: "STRING3", expectErr: true},
		{value: "TESTSMATESTSMATESTSMA", typeName: "STRING5", expectErr: true},
		{value: true, typeName: "STRING1", expectErr: true},
		{value: 123, typeName: "STRING3", expectErr: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ValidateDecimals(%v, %s)", test.value, test.typeName), func(t *testing.T) {
			result, err := validate.ValidateStrings(test.value, test.typeName)
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

package validate_test

import (
	"cbam_api/validate"
	"fmt"
	"testing"
)

func TestValidateNumerics(t *testing.T) {
	type testCase struct {
		value        interface{}
		typeName     string
		expectedPass bool
	}

	tests := []testCase{
		{value: 123, typeName: "NUMERIC3", expectedPass: true},
		{value: 1234567891, typeName: "NUMERIC10", expectedPass: true},
		{value: 1234, typeName: "NUMERIC3", expectedPass: false},
		{value: "not an int", typeName: "NUMERIC3", expectedPass: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ValidateNumerics(%v, %s)", test.value, test.typeName), func(t *testing.T) {
			pass, err := validate.ValidateNumerics(test.value, test.typeName)

			if !test.expectedPass {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if pass != test.expectedPass {
				t.Errorf("expected pass to be %t but got %t", test.expectedPass, pass)
			}
		})
	}
}

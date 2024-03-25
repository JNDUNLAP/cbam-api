package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateDecimals(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Float64 {
		return false, fmt.Errorf("expected float64, got %T", value)
	}
	v := rv.Float()

	valueStr := FormatFloatToString(v)

	totalDigits, err := ExtractTotalDigits("DECIMAL", typeName)
	if err != nil {
		return false, err
	}

	digitCount := len(strings.Replace(valueStr, ".", "", 1))
	if digitCount > totalDigits {
		return false, fmt.Errorf("expected up to %d total digits, got %d", totalDigits, digitCount)
	}
	return true, nil
}

func FormatFloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

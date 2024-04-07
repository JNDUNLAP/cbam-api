package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateNumerics(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Int {
		return false, fmt.Errorf("expected int, got %T", value)
	}

	v := rv.Int()

	totalDigits, err := strconv.Atoi(strings.TrimPrefix(typeName, "NUMERIC"))
	if err != nil {
		return false, err
	}

	if countDigits(v) != totalDigits {
		return false, fmt.Errorf("expected %d digits, got %d", totalDigits, countDigits(v))
	}

	return true, nil
}

func countDigits(n int64) int {
	count := 0
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

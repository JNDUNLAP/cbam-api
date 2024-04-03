package validate

import (
	"fmt"
	"reflect"
)

func ValidateStrings(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.String {
		return false, fmt.Errorf("expected string, got %T", value)
	}
	v := rv.String()
	minLength, maxLength, err := ExtractStringLengths(typeName)
	if err != nil {
		return false, err
	}
	length := len(v)
	if length < minLength || length > maxLength {
		return false, fmt.Errorf("expected String length between %d and %d characters, got %d", minLength, maxLength, length)
	}
	return true, nil
}

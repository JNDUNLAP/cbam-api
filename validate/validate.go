package validate

import (
	"cbam_api/files"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Validator func(value interface{}, typeName string) (bool, error)

func ExtractTotalDigits(typestem, typeName string) (int, error) {
	if !strings.HasPrefix(typeName, typestem) {
		return 0, fmt.Errorf("typeName does not start with '%s'", typestem)
	}
	numberStr := typeName[len(typestem):]
	totalDigits, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}
	return totalDigits, nil
}

func ExtractStringLengths(typeName string) (int, int, error) {
	if !strings.HasPrefix(typeName, "STRING") {
		return 0, 0, fmt.Errorf("type name does not start with 'AN'")
	}
	lengthStr := typeName[6:]
	maxLength, err := strconv.Atoi(lengthStr)
	if err != nil {
		return 0, 0, err
	}
	return 1, maxLength, nil
}

func validateField(fieldValue interface{}, fieldType string, validators map[string]Validator) error {
	for prefix, validator := range validators {
		if strings.HasPrefix(fieldType, prefix) {
			valid, err := validator(fieldValue, fieldType)
			if !valid {
				return fmt.Errorf("validation failed: %v", err)
			}
			break
		}
	}
	return nil
}
func validateStructRecursive(v reflect.Value, validators map[string]Validator) error {
	v = reflect.Indirect(v)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %s", v.Kind())
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanInterface() {
			fieldType := t.Field(i).Type
			fieldValue := field.Interface()

			if fieldType.Kind() == reflect.Struct {
				if err := validateStructRecursive(field, validators); err != nil {
					return err
				}
			} else if fieldType.Kind() == reflect.Slice && fieldType.Elem().Kind() == reflect.Struct {
				for j := 0; j < field.Len(); j++ {
					elem := field.Index(j)
					if err := validateStructRecursive(elem, validators); err != nil {
						return err
					}
				}
			} else {
				if err := validateField(fieldValue, fieldType.Name(), validators); err != nil {
					return fmt.Errorf("field '%s' failed validation: %v", t.Field(i).Name, err)
				}
			}
		}
	}

	return nil
}

func ValidateStruct(s interface{}, validators map[string]Validator) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct && v.Kind() != reflect.Ptr {
		return errors.New("ValidateStruct expects a struct or a pointer to a struct")
	}
	return validateStructRecursive(v, validators)
}

func DataModelValidate() error {
	fmt.Println("Testing Data Model Against EU Reporting Standards...")
	report, err := files.GetReport()
	if err != nil {
		return err // instead of printing and returning
	}

	validators := map[string]Validator{
		"DECIMAL": ValidateDecimals,
		"NUMERIC": ValidateNumerics,
		"STRING":  ValidateStrings,
	}

	err = ValidateStruct(report, validators)
	if err != nil {
		return fmt.Errorf("failed details %s", err)
	}
	return nil
}

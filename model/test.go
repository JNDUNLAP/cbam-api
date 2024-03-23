package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Validator func(value interface{}, typeName string) (bool, error)

func ValidateDateTime(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Struct || rv.Type() != reflect.TypeOf(time.Time{}) {
		return false, fmt.Errorf("expected time.Time, got %T", value)
	}
	dt := value.(time.Time)
	if !isValidDateTimeFormat(dt, "2006-01-02T15:04:05Z") {
		return false, fmt.Errorf("invalid DateTime format")
	}
	return true, nil
}

func isValidDateTimeFormat(dt time.Time, layout string) bool {
	return dt.UTC().Format(layout) == dt.Format(layout)
}

func ValidateDecimals(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Float64 {
		return false, fmt.Errorf("expected float64, got %T", value)
	}
	v := rv.Float()

	valueStr := formatFloatToString(v)

	totalDigits, err := extractTotalDigitsFromTypeName(typeName)
	if err != nil {
		return false, err
	}

	digitCount := len(strings.Replace(valueStr, ".", "", 1))
	if digitCount >= totalDigits {
		return false, fmt.Errorf("expected up to %d total digits, got %d", totalDigits, digitCount)
	}
	return true, nil
}

func formatFloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func extractTotalDigitsFromTypeName(typeName string) (int, error) {
	var totalDigits int
	_, err := fmt.Sscanf(typeName, "DECIMAL%d", &totalDigits)
	if err != nil {
		return 0, fmt.Errorf("failed to parse total digits from type name: %v", err)
	}
	return totalDigits, nil
}

func ValidateNumerics(value interface{}, typeName string) (bool, error) {
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Int {
		return false, fmt.Errorf("expected int, got %T", value)
	}
	v := rv.Int()
	totalDigits, err := extractTotalDigits(typeName)
	if err != nil {
		return false, err
	}
	valueStr := strconv.FormatInt(v, 10)
	if len(valueStr) != totalDigits {
		return false, fmt.Errorf("expected %d digits, got %d", totalDigits, len(valueStr))
	}
	return true, nil
}
func extractTotalDigits(typeName string) (int, error) {
	if !strings.HasPrefix(typeName, "NUMERIC") {
		return 0, errors.New("typeName does not start with 'NUMERIC'")
	}
	numberStr := typeName[7:]

	totalDigits, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}
	return totalDigits, nil
}

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

func GetReport() (QReport, error) {
	const dirPath = "files/xml/report/"
	const fileName = "Sample_CBAM_Quarterly_Report.xml"

	file, err := os.Open(filepath.Join(dirPath, fileName))
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return QReport{}, err
	}
	defer file.Close()

	report, err := UnmarshalReport(file)
	if err != nil {
		fmt.Printf("UnmarshalReport failed: %v\n", err)
		return QReport{}, err
	}

	return report, nil
}

func UnmarshalReport(file io.Reader) (QReport, error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return QReport{}, fmt.Errorf("failed to read file contents: %w", err)
	}

	var report QReport
	if err := xml.Unmarshal(fileBytes, &report); err != nil {
		return QReport{}, fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	return report, nil
}

func DataModelTest() {
	fmt.Println("\nTesting Data Model Against EU Standards...")
	report, err := GetReport()
	if err != nil {
		fmt.Println(err)
		return
	}

	validators := map[string]Validator{
		"DECIMAL": ValidateDecimals,
		"NUMERIC": ValidateNumerics,
		"STRING":  ValidateStrings,
	}

	err = ValidateStruct(report, validators)
	if err != nil {
		log.Fatal("Data Model: FAILED: ", err)
	} else {
		fmt.Println("Data Model: PASSED")
	}
}

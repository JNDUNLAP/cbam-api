package model

// import (
// 	"fmt"
// 	"reflect"
// )

// func GenerateErrorListForNonNullFields(v interface{}) ([]string, error) {
// 	var errors []string
// 	err := inspectStructForNonNullErrors(reflect.ValueOf(v), &errors, "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return errors, nil
// }

// func inspectStructForNonNullErrors(v reflect.Value, errors *[]string, path string) error {
// 	if v.Kind() == reflect.Ptr && !v.IsNil() {
// 		v = v.Elem()
// 	}

// 	if v.Kind() == reflect.Struct {
// 		t := v.Type()
// 		for i := 0; i < v.NumField(); i++ {
// 			field := v.Field(i)
// 			structField := t.Field(i)

// 			// Construct a path for nested fields
// 			currentPath := path
// 			if currentPath != "" {
// 				currentPath += "."
// 			}
// 			currentPath += structField.Name

// 			// Check if this field has an ErrorDetail and it's non-empty
// 			if structField.Type.Kind() == reflect.String && structField.Name == "ErrorDetail" && field.String() != "" {
// 				*errors = append(*errors, fmt.Sprintf("%s: %v", currentPath, field.String()))
// 			} else {
// 				// Recursive call for nested structs, slices, arrays
// 				inspectStructForNonNullErrors(field, errors, currentPath)
// 			}
// 		}
// 	} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
// 		for i := 0; i < v.Len(); i++ {
// 			inspectStructForNonNullErrors(v.Index(i), errors, fmt.Sprintf("%s[%d]", path, i))
// 		}
// 	}
// 	return nil
// }

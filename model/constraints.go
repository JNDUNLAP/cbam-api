package model

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var MarshalErrorsOnly bool = false

func SetupConstraints(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("SetupConstraints expects a non-nil pointer to a struct, got %T", v)
	}
	log.Println("Starting to apply constraints...")
	return applyConstraints(val.Elem(), "")
}

func applyConstraints(val reflect.Value, path string) error {
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)
		fieldPath := fmt.Sprintf("%s.%s", path, structField.Name)
		if err := processField(field, structField, fieldPath); err != nil {
			return err
		}
	}
	return nil
}

func processField(field reflect.Value, structField reflect.StructField, path string) error {
	if structField.PkgPath != "" {
		return nil
	}

	if field.Kind() == reflect.Ptr && !field.IsNil() {
		field = field.Elem()
	}

	if isCustomType(field) {
		return applyConstraintsToField(field, structField, path)
	}

	switch field.Kind() {
	case reflect.Struct:
		return applyConstraints(field, path)
	case reflect.Slice, reflect.Array:
		return processSliceOrArray(field, path)
	}

	return nil
}

func isCustomType(field reflect.Value) bool {
	if _, ok := field.Interface().(ConstraintSetter); ok {
		return true
	}
	if field.CanAddr() {
		_, ok := field.Addr().Interface().(ConstraintSetter)
		return ok
	}
	return false
}

func applyConstraintsToField(field reflect.Value, structField reflect.StructField, path string) error {
	constraints := extractConstraints(structField)

	// Initialize nil pointer fields to their zero value if needed
	if field.Kind() == reflect.Ptr && field.IsNil() {
		newInstance := reflect.New(structField.Type.Elem())
		field.Set(newInstance)
		field = newInstance.Elem() // Update field to reference the new instance
	}

	// Check if the field implements ConstraintSetter and apply constraints
	if setter, ok := extractSetter(field); ok {
		if err := setter.ApplyConstraints(constraints); err != nil {
			log.Printf("Error applying constraints to field '%s': %v\n", path, err)
			return err
		}
		if err := setter.ValidateConstraints(); err != nil {
			log.Printf("Validation error for field '%s': %v\n", path, err)
			return err
		}
	}

	return nil
}
func extractSetter(field reflect.Value) (ConstraintSetter, bool) {
	if field.CanAddr() && field.Addr().CanInterface() {
		if setter, ok := field.Addr().Interface().(ConstraintSetter); ok {
			return setter, true
		}
	}
	if field.CanInterface() {
		if setter, ok := field.Interface().(ConstraintSetter); ok {
			return setter, true
		}
	}
	return nil, false
}
func extractConstraints(field reflect.StructField) map[string]string {
	constraints := make(map[string]string)
	for _, key := range []string{"min", "max", "totalDigits", "fracDigits"} {
		if value, ok := field.Tag.Lookup(key); ok {
			constraints[key] = value
		}
	}
	return constraints
}

func processSliceOrArray(field reflect.Value, path string) error {
	for i := 0; i < field.Len(); i++ {
		elem := field.Index(i)
		if elem.Kind() == reflect.Ptr && !elem.IsNil() {
			elem = elem.Elem()
		}
		if elem.Kind() == reflect.Struct {
			elemPath := fmt.Sprintf("%s[%d]", path, i)
			if err := applyConstraints(elem, elemPath); err != nil {
				return err
			}
		}
	}
	return nil
}

type ConstraintSetter interface {
	ApplyConstraints(constraints map[string]string) error
	ValidateConstraints() error
}

func (c *ConstrainedString) SetConstraints(min, max int) {
	c.Min = min
	c.Max = max
}

func (c *ConstrainedInt) SetConstraints(min, max int) {
	c.Min = min
	c.Max = max
}

func (c *ConstrainedString) ValidateConstraints() error {
	if c.Min != 0 || c.Max != 0 {
		valueLength := len(c.Value)
		if valueLength < c.Min || valueLength > c.Max {
			err := fmt.Errorf("ERROR: length %d is out of bounds [%d, %d]", valueLength, c.Min, c.Max)
			c.ErrorDetail = err.Error()
			return err
		}
	}
	c.ErrorDetail = ""
	return nil
}

func (c *ConstrainedInt) ValidateConstraints() error {
	valueStr := strconv.Itoa(c.Value)
	valueLength := len(valueStr)
	if valueLength < c.Min || valueLength > c.Max {
		err := fmt.Errorf("ERROR: length %d is out of bounds [%d, %d]", valueLength, c.Min, c.Max)
		c.ErrorDetail = err.Error()
		return err
	}
	c.ErrorDetail = ""
	return nil
}

func (c *ConstrainedDecimal) ValidateConstraints() error {
	if c.Value == nil {
		return fmt.Errorf("null")
	}

	// Convert to string to check total and fractional digits.
	valStr := c.Value.Text('f', -1) // Get string representation with full precision.
	parts := strings.Split(valStr, ".")
	totalDigits := len(parts[0])
	fracDigits := 0
	if len(parts) > 1 {
		fracDigits = len(parts[1])
	}

	// Validate against the totalDigits and fracDigits constraints.
	if totalDigits > c.TotalDigits || fracDigits > c.FracDigits {
		c.ErrorDetail = fmt.Sprintf("value %s exceeds allowed digits (total: %d, fractional: %d)", valStr, c.TotalDigits, c.FracDigits)
		return fmt.Errorf(c.ErrorDetail)
	}

	return nil
}

func (c *ConstrainedDecimal) ApplyConstraints(constraints map[string]string) {
	if totalDigits, ok := constraints["totalDigits"]; ok {
		if td, err := strconv.Atoi(totalDigits); err == nil {
			c.TotalDigits = td
		}
	}
	if fracDigits, ok := constraints["fracDigits"]; ok {
		if fd, err := strconv.Atoi(fracDigits); err == nil {
			c.FracDigits = fd
		}
	}
}

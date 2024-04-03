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
	xmlTagName := structField.Tag.Get("xml")
	xmlFieldName := strings.Split(xmlTagName, ",")[0]

	minTag := structField.Tag.Get("min")
	maxTag := structField.Tag.Get("max")

	var min, max int
	var err error

	if minTag != "" {
		min, err = strconv.Atoi(minTag)
		if err != nil {
			log.Printf("Warning: Invalid 'min' constraint for XML field '%s' (Go field '%s'). Error: %v\n", xmlFieldName, structField.Name, err)
		}
	}
	if maxTag != "" {
		max, err = strconv.Atoi(maxTag)
		if err != nil {
			log.Printf("Warning: Invalid 'max' constraint for XML field '%s' (Go field '%s'). Error: %v\n", xmlFieldName, structField.Name, err)
		}
	}

	if field.Kind() == reflect.Ptr && field.IsNil() {
		if structField.Type.Elem() == reflect.TypeOf(ConstrainedString{}) {
			field.Set(reflect.New(structField.Type.Elem()))
		}
	}

	if field.CanAddr() && field.Addr().CanInterface() {
		if setter, ok := field.Addr().Interface().(ConstraintSetter); ok {
			// log.Printf("Applying constraints [Min: %d, Max: %d] to XML field '%s' ('%s')\n", min, max, xmlFieldName, structField.Name)
			setter.SetConstraints(min, max)
			setter.ValidateConstraints()
		} else if field.CanInterface() {
			if setter, ok := field.Interface().(ConstraintSetter); ok {
				// log.Printf("Applying constraints [Min: %d, Max: %d] to XML field '%s' ('%s')\n", min, max, xmlFieldName, structField.Name)
				setter.SetConstraints(min, max)
				setter.ValidateConstraints()
			}
		}
	}

	return nil
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
	SetConstraints(min, max int)
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

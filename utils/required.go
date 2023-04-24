package utils

import (
	"fmt"
	"reflect"
)

func isNilSafe(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Slice, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func CheckRequiredFieldAllSet(object interface{}) error {
	v := reflect.ValueOf(object)
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("struct input required")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)

		requiredTag := fieldType.Tag.Get("required")

		if requiredTag != "true" {
			continue
		}

		fieldValue := v.Field(i)

		if isNilSafe(fieldValue) {
			return fmt.Errorf("field `%s` is required and must be specified in %s", fieldType.Name, t.Name())
		}
	}

	return nil
}

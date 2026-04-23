package utils

import (
	"reflect"
)

func CheckImportStringBlankOrNull(params ...any) bool {
	for _, param := range params {
		if IsValueBlank(param) {
			return true
		}
	}
	return false
}

func IsValueBlank(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if rv.IsNil() {
			return true
		}
	}
	return reflect.DeepEqual(value, reflect.Zero(rv.Type()).Interface())
}

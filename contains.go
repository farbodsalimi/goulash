package goulash

import (
	"fmt"
	"reflect"
	"strings"
)

func Contains[T, M any](object T, element M) bool {
	objVal := reflect.ValueOf(object)
	elemStr := fmt.Sprint(element)

	switch objVal.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < objVal.Len(); i++ {
			if fmt.Sprint(objVal.Index(i).Interface()) == elemStr {
				return true
			}
		}

	case reflect.Map:
		for _, key := range objVal.MapKeys() {
			if fmt.Sprint(key.Interface()) == elemStr {
				return true
			}
		}

	case reflect.String:
		if elem, ok := any(element).(string); ok {
			return strings.Contains(objVal.String(), elem)
		}
	}

	return false
}

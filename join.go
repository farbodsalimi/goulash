package goulash

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/farbodsalimi/goulash/utils"
)

// Join creates and returns a new string by concatenating all of the elements in a slice, separated by a specified delimiter
func Join[T any](slice []T, delimiter string) string {
	var result string

	generic := utils.ParseGeneric(slice)
	switch generic.Type.Elem().Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		result = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delimiter), "[]")

	case reflect.String:
		result = strings.Join(generic.Value.Interface().([]string), delimiter)
	}

	return result
}

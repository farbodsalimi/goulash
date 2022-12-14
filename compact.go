package goulash

import (
	"reflect"

	"github.com/farbodsalimi/goulash/utils"
	"golang.org/x/exp/constraints"
)

func Compact[T constraints.Ordered](slice []T) []T {
	result := []T{}

	for _, value := range slice {
		parsedGeneric := utils.ParseGeneric(value)

		switch parsedGeneric.Kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if parsedGeneric.Value.Int() != 0 {
				result = append(result, value)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if parsedGeneric.Value.Uint() != 0 {
				result = append(result, value)
			}
		case reflect.Float32, reflect.Float64:
			if parsedGeneric.Value.Float() != 0 {
				result = append(result, value)
			}
		case reflect.String:
			s := parsedGeneric.Value.String()
			if s != "" && len(s) != 0 {
				result = append(result, value)
			}
		}
	}

	return result
}

package goulash

import (
	"reflect"

	"github.com/farbodsalimi/goulash/utils"
	"golang.org/x/exp/constraints"
)

func Product[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	parsedFirst := utils.ParseGeneric(slice[0])
	switch parsedFirst.Kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result := int64(1)
		for _, item := range slice {
			parsed := utils.ParseGeneric(item)
			result *= parsed.Value.Int()
		}
		convertedResult := reflect.ValueOf(result).Convert(parsedFirst.Type)
		return convertedResult.Interface().(T)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := uint64(1)
		for _, item := range slice {
			parsed := utils.ParseGeneric(item)
			result *= parsed.Value.Uint()
		}
		convertedResult := reflect.ValueOf(result).Convert(parsedFirst.Type)
		return convertedResult.Interface().(T)
	case reflect.Float32, reflect.Float64:
		result := float64(1)
		for _, item := range slice {
			parsed := utils.ParseGeneric(item)
			result *= parsed.Value.Float()
		}
		convertedResult := reflect.ValueOf(result).Convert(parsedFirst.Type)
		return convertedResult.Interface().(T)
	}
	var zero T
	return zero
}

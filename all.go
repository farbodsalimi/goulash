package goulash

import (
	"reflect"

	"github.com/farbodsalimi/goulash/utils"
	"golang.org/x/exp/constraints"
)

func All[T constraints.Ordered](slice []T) bool {
	for _, value := range slice {
		parsedGeneric := utils.ParseGeneric(value)

		switch parsedGeneric.Kind {
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
			if parsedGeneric.Value.IsZero() {
				return false
			}

		case reflect.String:
			s := parsedGeneric.Value.String()
			if s == "" || len(s) == 0 {
				return false
			}
		}
	}

	return true
}

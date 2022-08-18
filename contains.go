package goulash

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/farbodsalimi/goulash/utils"
)

func Contains[T, M any](object T, element M) bool {
	parsedObject := utils.ParseGeneric(object)
	parsedElement := utils.ParseGeneric(element)

	switch parsedObject.Kind {

	case reflect.Slice, reflect.Array:
		for i := 0; i < parsedObject.Value.Len(); i++ {
			if fmt.Sprint(parsedObject.Value.Index(i)) == fmt.Sprint(parsedElement.Value) {
				return true
			}
		}

	case reflect.Map:
		for _, key := range parsedObject.Value.MapKeys() {
			if fmt.Sprint(key) == fmt.Sprint(parsedElement.Value) {
				return true
			}
		}

	case reflect.String:
		return strings.Contains(parsedObject.Value.String(), parsedElement.Value.String())

	default:
		return false
	}

	return false
}

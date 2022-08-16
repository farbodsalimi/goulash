package utils

import "reflect"

func TypeOfGeneric(value any) reflect.Kind {
	return reflect.ValueOf(value).Type().Elem().Kind()
}

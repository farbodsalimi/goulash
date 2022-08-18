package utils

import "reflect"

type Generic struct {
	Value reflect.Value
	Type  reflect.Type
	Kind  reflect.Kind
}

func ParseGeneric(generic any) (g *Generic) {
	value := reflect.ValueOf(generic)
	kind := value.Kind()
	gtype := value.Type()

	return &Generic{
		Value: value,
		Kind:  kind,
		Type:  gtype,
	}
}

package utils

import "reflect"

type Generic struct {
	Value reflect.Value
	Type  reflect.Type
	Kind  reflect.Kind
}

func ParseGeneric(generic any) (g *Generic) {
	g = &Generic{
		Value: reflect.ValueOf(generic),
	}

	if !g.Value.IsZero() {
		g.Type = g.Value.Type()
		g.Kind = g.Value.Kind()
	}

	return g
}

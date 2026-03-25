package goulash

import "golang.org/x/exp/constraints"

func Product[T constraints.Integer | constraints.Float](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	result := T(1)
	for _, item := range slice {
		result *= item
	}
	return result
}

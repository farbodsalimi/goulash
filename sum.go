package goulash

import "golang.org/x/exp/constraints"

func Sum[T constraints.Ordered](slice []T) T {
	var sum T
	for _, item := range slice {
		sum += item
	}
	return sum
}

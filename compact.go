package goulash

import "golang.org/x/exp/constraints"

func Compact[T constraints.Ordered](slice []T) []T {
	var zero T
	result := []T{}
	for _, value := range slice {
		if value != zero {
			result = append(result, value)
		}
	}
	return result
}

package goulash

import "golang.org/x/exp/constraints"

// All returns true only if every element in slice is "truthy" (non-zero).
// A zero value of T is considered falsy (0 for numbers, "" for strings).
func All[T constraints.Ordered](slice []T) bool {
	var zero T // zero value for T: 0 for numeric types, "" for strings
	for _, value := range slice {
		if value == zero {
			return false
		}
	}
	return true
}

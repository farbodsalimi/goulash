package goulash

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](slice []T) []T {
	toBeSorted := make([]T, len(slice))
	copy(toBeSorted, slice)

	slices.Sort(toBeSorted)

	return toBeSorted
}

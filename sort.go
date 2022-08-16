package goulash

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](slice []T) []T {
	toBeSorted := make([]T, len(slice))
	copy(toBeSorted, slice)

	sort.Slice(toBeSorted, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return toBeSorted
}

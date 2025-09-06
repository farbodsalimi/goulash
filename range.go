package goulash

import "golang.org/x/exp/constraints"

func Range[T constraints.Integer](start, end T) []T {
	if start >= end {
		return []T{}
	}

	result := make([]T, 0, int(end-start))
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

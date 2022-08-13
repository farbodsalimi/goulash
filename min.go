package goulash

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}

	return min
}

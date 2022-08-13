package goulash

import "golang.org/x/exp/constraints"

func MinMax[T constraints.Ordered](numbers ...T) (min T, max T) {
	return Min(numbers...), Max(numbers...)
}

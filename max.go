package goulash

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return max
}

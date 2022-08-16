package goulash

import "golang.org/x/exp/constraints"

func GroupBy[T constraints.Ordered, M constraints.Ordered](
	slice []T,
	fn func(currentValue T) M,
) map[M][]T {
	grouped := make(map[M][]T, len(slice))
	for _, value := range slice {
		key := fn(value)
		if _, ok := grouped[key]; ok {
			grouped[key] = append(grouped[key], value)
		} else {
			grouped[key] = []T{value}
		}
	}
	return grouped
}

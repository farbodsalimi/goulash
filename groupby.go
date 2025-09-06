package goulash

import "golang.org/x/exp/constraints"

func GroupBy[T constraints.Ordered, M constraints.Ordered](
	slice []T,
	fn func(currentValue T) M,
) map[M][]T {
	grouped := make(map[M][]T, len(slice))
	for _, value := range slice {
		key := fn(value)
		grouped[key] = append(grouped[key], value)
	}
	return grouped
}

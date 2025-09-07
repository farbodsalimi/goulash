package goulash

func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var truthy, falsy []T
	for _, item := range slice {
		if predicate(item) {
			truthy = append(truthy, item)
		} else {
			falsy = append(falsy, item)
		}
	}
	return truthy, falsy
}

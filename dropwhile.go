package goulash

func DropWhile[T any](slice []T, predicate func(T) bool) []T {
	for i, item := range slice {
		if !predicate(item) {
			return append([]T{}, slice[i:]...)
		}
	}
	return []T{}
}

package goulash

func None[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return false
		}
	}
	return true
}

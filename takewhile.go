package goulash

func TakeWhile[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if !predicate(item) {
			break
		}
		result = append(result, item)
	}
	return result
}

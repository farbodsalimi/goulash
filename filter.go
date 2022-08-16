package goulash

func Filter[T any](slice []T, fn func(T) bool) []T {
	var filtered []T
	for _, value := range slice {
		if fn(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

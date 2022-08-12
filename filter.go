package main

func Filter[T any](slice []T, fn func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

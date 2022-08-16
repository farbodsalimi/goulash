package goulash

func ForEach[T any](slice []T, fn func(currentValue T, args ...any)) {
	for index, value := range slice {
		fn(value, index)
	}
}

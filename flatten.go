package goulash

func Flatten[T any](nested [][]T) []T {
	var result []T
	for _, slice := range nested {
		result = append(result, slice...)
	}
	return result
}

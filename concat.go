package goulash

func Concat[T any](args ...[]T) []T {
	var result []T

	for _, v := range args {
		result = append(result, v...)
	}

	return result
}

package goulash

func Concat[T any](slices ...[]T) []T {
	result := []T{}
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

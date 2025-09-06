package goulash

func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return append([]T{}, slice...)
	}
	if n >= len(slice) {
		return []T{}
	}
	return append([]T{}, slice[n:]...)
}

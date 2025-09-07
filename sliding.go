package goulash

func Sliding[T any](slice []T, size int) [][]T {
	if size <= 0 || len(slice) == 0 {
		return [][]T{}
	}

	if size > len(slice) {
		return [][]T{append([]T{}, slice...)}
	}

	var result [][]T
	for i := 0; i <= len(slice)-size; i++ {
		window := append([]T{}, slice[i:i+size]...)
		result = append(result, window)
	}
	return result
}

package goulash

func Repeat[T any](item T, count int) []T {
	if count <= 0 {
		return []T{}
	}

	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = item
	}
	return result
}

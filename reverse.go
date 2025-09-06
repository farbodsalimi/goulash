package goulash

func Reverse[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	result := make([]T, len(slice))
	for i, item := range slice {
		result[len(slice)-1-i] = item
	}
	return result
}

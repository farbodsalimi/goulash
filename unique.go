package goulash

func Unique[T comparable](slice []T) []T {
	lookup := make(map[T]struct{})
	result := []T{}
	for _, value := range slice {
		if _, ok := lookup[value]; !ok {
			lookup[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

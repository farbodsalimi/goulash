package goulash

func Unique[T comparable](slice []T) []T {
	lookup := make(map[T]bool)
	result := []T{}
	for _, value := range slice {
		if _, ok := lookup[value]; !ok {
			lookup[value] = true
			result = append(result, value)
		}
	}
	return result
}

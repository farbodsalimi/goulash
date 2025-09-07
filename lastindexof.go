package goulash

func LastIndexOf[T comparable](slice []T, element T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == element {
			return i
		}
	}
	return -1
}

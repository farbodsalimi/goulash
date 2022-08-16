package goulash

func Union[T comparable](slices ...[]T) []T {
	var result []T
	lookup := make(map[T]bool)

	for _, slice := range slices {
		for _, value := range slice {
			if _, ok := lookup[value]; !ok {
				result = append(result, value)
				lookup[value] = true
			}
		}
	}

	return result
}

package goulash

func Union[T comparable](slices ...[]T) []T {
	var result []T
	lookup := make(map[T]bool)

	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := lookup[item]; !ok {
				result = append(result, item)
				lookup[item] = true
			}
		}
	}

	return result
}

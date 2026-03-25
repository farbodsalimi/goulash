package goulash

func Union[T comparable](slices ...[]T) []T {
	result := []T{}
	lookup := make(map[T]struct{})

	for _, slice := range slices {
		for _, value := range slice {
			if _, ok := lookup[value]; !ok {
				result = append(result, value)
				lookup[value] = struct{}{}
			}
		}
	}

	return result
}

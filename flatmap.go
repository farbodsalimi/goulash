package goulash

func FlatMap[T any, U any](ts []T, fn func(T) []U) []U {
	var result []U

	for _, t := range ts {
		mappedValues := fn(t)
		result = append(result, mappedValues...)
	}

	return result
}

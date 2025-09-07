package goulash

func Permutations[T any](slice []T) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}

	if len(slice) == 1 {
		return [][]T{{slice[0]}}
	}

	var result [][]T
	for i, item := range slice {
		remaining := append(append([]T{}, slice[:i]...), slice[i+1:]...)
		subPerms := Permutations(remaining)
		for _, perm := range subPerms {
			newPerm := append([]T{item}, perm...)
			result = append(result, newPerm)
		}
	}

	return result
}

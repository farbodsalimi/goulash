package goulash

func Combinations[T any](slice []T, r int) [][]T {
	if r <= 0 || r > len(slice) {
		return [][]T{}
	}

	if r == len(slice) {
		return [][]T{append([]T{}, slice...)}
	}

	if r == 1 {
		var result [][]T
		for _, item := range slice {
			result = append(result, []T{item})
		}
		return result
	}

	var result [][]T
	for i := 0; i <= len(slice)-r; i++ {
		item := slice[i]
		remaining := slice[i+1:]
		subCombinations := Combinations(remaining, r-1)
		for _, comb := range subCombinations {
			newComb := append([]T{item}, comb...)
			result = append(result, newComb)
		}
	}

	return result
}

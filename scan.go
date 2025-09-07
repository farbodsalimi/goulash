package goulash

func Scan[T, M any](slice []T, fn func(M, T) M, initValue M) []M {
	if len(slice) == 0 {
		return []M{initValue}
	}

	result := make([]M, len(slice)+1)
	result[0] = initValue

	acc := initValue
	for i, item := range slice {
		acc = fn(acc, item)
		result[i+1] = acc
	}

	return result
}

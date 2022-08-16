package goulash

func Map[T any, M any](slice []T, fn func(currentValue T) M) []M {
	n := make([]M, len(slice))
	for index, value := range slice {
		n[index] = fn(value)
	}
	return n
}

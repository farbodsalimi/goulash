package goulash

func Reduce[T, M any](slice []T, fn func(M, T) M, initValue M) M {
	accumulator := initValue
	for _, value := range slice {
		accumulator = fn(accumulator, value)
	}
	return accumulator
}

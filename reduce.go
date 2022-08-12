package main

func Reduce[T, M any](slice []T, fn func(M, T) M, initValue M) M {
	accumulator := initValue
	for _, v := range slice {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

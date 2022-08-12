package main

func Concat[T any](slice1 []T, slice2 []T) []T {
	return append(slice1, slice2...)
}

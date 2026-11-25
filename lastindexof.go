package goulash

import "slices"

func LastIndexOf[T comparable](slice []T, element T) int {
	for i, s := range slices.Backward(slice) {
		if s == element {
			return i
		}
	}
	return -1
}

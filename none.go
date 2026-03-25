package goulash

import "slices"

func None[T any](slice []T, predicate func(T) bool) bool {
	return !slices.ContainsFunc(slice, predicate)
}

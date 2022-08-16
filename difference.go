package goulash

import "golang.org/x/exp/constraints"

func Difference[T constraints.Ordered](slice1 []T, slice2 []T) []T {
	lookup := map[T]int{}

	for _, s := range slice1 {
		lookup[s]++
	}
	for _, s := range slice2 {
		lookup[s]--
	}

	result := []T{}
	for s, v := range lookup {
		if v != 0 {
			result = append(result, s)
		}
	}

	return result
}

package goulash

import "golang.org/x/exp/constraints"

func Difference[T constraints.Ordered](slice1 []T, slice2 []T) []T {
	lookup := map[T]int{}

	for _, value := range slice1 {
		lookup[value]++
	}
	for _, value := range slice2 {
		lookup[value]--
	}

	result := []T{}
	for s, value := range lookup {
		if value != 0 {
			result = append(result, s)
		}
	}

	return result
}

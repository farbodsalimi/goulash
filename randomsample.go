package goulash

import (
	"math/rand"
	"time"
)

func RandomSample[T any](slice []T, count int) []T {
	if count <= 0 || len(slice) == 0 {
		return []T{}
	}

	if count >= len(slice) {
		return append([]T{}, slice...)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	indices := r.Perm(len(slice))
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = slice[indices[i]]
	}

	return result
}

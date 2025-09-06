package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomSample(t *testing.T) {
	t.Run("RandomSampleNormal", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := RandomSample(numbers, 3)

		assert.Len(t, result, 3)
		// Check that all sampled elements are from original slice
		for _, item := range result {
			assert.Contains(t, numbers, item)
		}
	})

	t.Run("RandomSampleMoreThanLength", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := RandomSample(numbers, 5)

		assert.Equal(t, numbers, result)
	})

	t.Run("RandomSampleZero", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RandomSample(numbers, 0)

		assert.Empty(t, result)
	})

	t.Run("RandomSampleEmpty", func(t *testing.T) {
		var empty []int
		result := RandomSample(empty, 3)

		assert.Empty(t, result)
	})

	t.Run("RandomSampleOne", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RandomSample(numbers, 1)

		assert.Len(t, result, 1)
		assert.Contains(t, numbers, result[0])
	})
}

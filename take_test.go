package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	t.Run("TakeNormalCase", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Take(numbers, 3)

		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("TakeZero", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Take(numbers, 0)

		assert.Empty(t, result)
	})

	t.Run("TakeNegative", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Take(numbers, -1)

		assert.Empty(t, result)
	})

	t.Run("TakeMoreThanLength", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Take(numbers, 5)

		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("TakeFromEmpty", func(t *testing.T) {
		var empty []int
		result := Take(empty, 3)

		assert.Empty(t, result)
	})

	t.Run("TakeStrings", func(t *testing.T) {
		words := []string{"hello", "world", "go", "programming"}
		result := Take(words, 2)

		assert.Equal(t, []string{"hello", "world"}, result)
	})
}

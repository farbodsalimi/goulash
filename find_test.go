package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("FindExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result, found := Find(numbers, func(n int) bool {
			return n > 3
		})

		assert.True(t, found)
		assert.Equal(t, 4, result)
	})

	t.Run("FindNotExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result, found := Find(numbers, func(n int) bool {
			return n > 10
		})

		assert.False(t, found)
		assert.Equal(t, 0, result)
	})

	t.Run("FindEmpty", func(t *testing.T) {
		var empty []int
		result, found := Find(empty, func(n int) bool {
			return n > 0
		})

		assert.False(t, found)
		assert.Equal(t, 0, result)
	})

	t.Run("FindStrings", func(t *testing.T) {
		words := []string{"hello", "world", "go"}
		result, found := Find(words, func(s string) bool {
			return len(s) == 2
		})

		assert.True(t, found)
		assert.Equal(t, "go", result)
	})
}

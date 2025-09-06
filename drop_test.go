package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop(t *testing.T) {
	t.Run("DropNormalCase", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Drop(numbers, 2)

		assert.Equal(t, []int{3, 4, 5}, result)
	})

	t.Run("DropZero", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Drop(numbers, 0)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("DropNegative", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Drop(numbers, -1)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("DropMoreThanLength", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Drop(numbers, 5)

		assert.Empty(t, result)
	})

	t.Run("DropFromEmpty", func(t *testing.T) {
		var empty []int
		result := Drop(empty, 3)

		assert.Empty(t, result)
	})

	t.Run("DropStrings", func(t *testing.T) {
		words := []string{"hello", "world", "go", "programming"}
		result := Drop(words, 1)

		assert.Equal(t, []string{"world", "go", "programming"}, result)
	})
}

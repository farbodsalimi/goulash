package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Run("IndexOfFound", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := IndexOf(numbers, 3)

		assert.Equal(t, 2, index)
	})

	t.Run("IndexOfNotFound", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := IndexOf(numbers, 10)

		assert.Equal(t, -1, index)
	})

	t.Run("IndexOfFirst", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2, 5}
		index := IndexOf(numbers, 2)

		assert.Equal(t, 1, index)
	})

	t.Run("IndexOfStrings", func(t *testing.T) {
		words := []string{"hello", "world", "go"}
		index := IndexOf(words, "world")

		assert.Equal(t, 1, index)
	})

	t.Run("IndexOfEmpty", func(t *testing.T) {
		var empty []int
		index := IndexOf(empty, 1)

		assert.Equal(t, -1, index)
	})
}

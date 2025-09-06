package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastIndexOf(t *testing.T) {
	t.Run("LastIndexOfFound", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2, 5}
		index := LastIndexOf(numbers, 2)

		assert.Equal(t, 3, index)
	})

	t.Run("LastIndexOfNotFound", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := LastIndexOf(numbers, 10)

		assert.Equal(t, -1, index)
	})

	t.Run("LastIndexOfSingle", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := LastIndexOf(numbers, 3)

		assert.Equal(t, 2, index)
	})

	t.Run("LastIndexOfStrings", func(t *testing.T) {
		words := []string{"hello", "world", "hello"}
		index := LastIndexOf(words, "hello")

		assert.Equal(t, 2, index)
	})

	t.Run("LastIndexOfEmpty", func(t *testing.T) {
		var empty []int
		index := LastIndexOf(empty, 1)

		assert.Equal(t, -1, index)
	})
}

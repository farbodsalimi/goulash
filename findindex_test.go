package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	t.Run("FindIndexExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := FindIndex(numbers, func(n int) bool {
			return n > 3
		})

		assert.Equal(t, 3, index)
	})

	t.Run("FindIndexNotExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := FindIndex(numbers, func(n int) bool {
			return n > 10
		})

		assert.Equal(t, -1, index)
	})

	t.Run("FindIndexEmpty", func(t *testing.T) {
		var empty []int
		index := FindIndex(empty, func(n int) bool {
			return n > 0
		})

		assert.Equal(t, -1, index)
	})

	t.Run("FindIndexFirst", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		index := FindIndex(numbers, func(n int) bool {
			return n > 0
		})

		assert.Equal(t, 0, index)
	})
}

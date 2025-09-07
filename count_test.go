package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	t.Run("CountMatching", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		count := Count(numbers, func(n int) bool {
			return n%2 == 0
		})

		assert.Equal(t, 3, count)
	})

	t.Run("CountNoneMatching", func(t *testing.T) {
		numbers := []int{1, 3, 5}
		count := Count(numbers, func(n int) bool {
			return n%2 == 0
		})

		assert.Equal(t, 0, count)
	})

	t.Run("CountAllMatching", func(t *testing.T) {
		numbers := []int{2, 4, 6}
		count := Count(numbers, func(n int) bool {
			return n%2 == 0
		})

		assert.Equal(t, 3, count)
	})

	t.Run("CountEmpty", func(t *testing.T) {
		var empty []int
		count := Count(empty, func(n int) bool {
			return n > 0
		})

		assert.Equal(t, 0, count)
	})
}

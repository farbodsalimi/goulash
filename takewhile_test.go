package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	t.Run("TakeWhileNumbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := TakeWhile(numbers, func(n int) bool {
			return n < 4
		})

		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("TakeWhileStrings", func(t *testing.T) {
		words := []string{"a", "ab", "abc", "abcd", "x"}
		result := TakeWhile(words, func(s string) bool {
			return len(s) <= 3
		})

		assert.Equal(t, []string{"a", "ab", "abc"}, result)
	})

	t.Run("TakeWhileEmpty", func(t *testing.T) {
		var empty []int
		result := TakeWhile(empty, func(n int) bool {
			return n > 0
		})

		assert.Empty(t, result)
	})

	t.Run("TakeWhileNoneMatch", func(t *testing.T) {
		numbers := []int{5, 6, 7, 8}
		result := TakeWhile(numbers, func(n int) bool {
			return n < 5
		})

		assert.Empty(t, result)
	})

	t.Run("TakeWhileAllMatch", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := TakeWhile(numbers, func(n int) bool {
			return n < 10
		})

		assert.Equal(t, []int{1, 2, 3, 4}, result)
	})
}

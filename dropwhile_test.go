package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropWhile(t *testing.T) {
	t.Run("DropWhileNumbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := DropWhile(numbers, func(n int) bool {
			return n < 4
		})

		assert.Equal(t, []int{4, 5, 6}, result)
	})

	t.Run("DropWhileStrings", func(t *testing.T) {
		words := []string{"a", "ab", "abc", "abcd", "x"}
		result := DropWhile(words, func(s string) bool {
			return len(s) <= 3
		})

		assert.Equal(t, []string{"abcd", "x"}, result)
	})

	t.Run("DropWhileEmpty", func(t *testing.T) {
		var empty []int
		result := DropWhile(empty, func(n int) bool {
			return n > 0
		})

		assert.Empty(t, result)
	})

	t.Run("DropWhileNoneMatch", func(t *testing.T) {
		numbers := []int{5, 6, 7, 8}
		result := DropWhile(numbers, func(n int) bool {
			return n < 5
		})

		assert.Equal(t, []int{5, 6, 7, 8}, result)
	})

	t.Run("DropWhileAllMatch", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := DropWhile(numbers, func(n int) bool {
			return n < 10
		})

		assert.Empty(t, result)
	})
}

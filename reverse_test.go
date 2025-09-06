package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("ReverseNumbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Reverse(numbers)

		assert.Equal(t, []int{5, 4, 3, 2, 1}, result)
	})

	t.Run("ReverseStrings", func(t *testing.T) {
		words := []string{"hello", "world", "go"}
		result := Reverse(words)

		assert.Equal(t, []string{"go", "world", "hello"}, result)
	})

	t.Run("ReverseEmpty", func(t *testing.T) {
		var empty []int
		result := Reverse(empty)

		assert.Empty(t, result)
	})

	t.Run("ReverseSingle", func(t *testing.T) {
		single := []int{42}
		result := Reverse(single)

		assert.Equal(t, []int{42}, result)
	})
}

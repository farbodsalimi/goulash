package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	t.Run("PermutationsThreeItems", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Permutations(slice)

		expected := [][]int{
			{1, 2, 3}, {1, 3, 2},
			{2, 1, 3}, {2, 3, 1},
			{3, 1, 2}, {3, 2, 1},
		}
		assert.Len(t, result, 6)
		for _, perm := range expected {
			assert.Contains(t, result, perm)
		}
	})

	t.Run("PermutationsTwoItems", func(t *testing.T) {
		slice := []string{"a", "b"}
		result := Permutations(slice)

		expected := [][]string{{"a", "b"}, {"b", "a"}}
		assert.Len(t, result, 2)
		for _, perm := range expected {
			assert.Contains(t, result, perm)
		}
	})

	t.Run("PermutationsOneItem", func(t *testing.T) {
		slice := []int{42}
		result := Permutations(slice)

		expected := [][]int{{42}}
		assert.Equal(t, expected, result)
	})

	t.Run("PermutationsEmpty", func(t *testing.T) {
		var empty []int
		result := Permutations(empty)

		assert.Empty(t, result)
	})
}

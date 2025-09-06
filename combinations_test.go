package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	t.Run("CombinationsThreeChooseTwo", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Combinations(slice, 2)

		expected := [][]int{{1, 2}, {1, 3}, {2, 3}}
		assert.Len(t, result, 3)
		for _, comb := range expected {
			assert.Contains(t, result, comb)
		}
	})

	t.Run("CombinationsFourChooseTwo", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		result := Combinations(slice, 2)

		expected := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
		assert.Len(t, result, 6)
		for _, comb := range expected {
			assert.Contains(t, result, comb)
		}
	})

	t.Run("CombinationsChooseOne", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		result := Combinations(slice, 1)

		expected := [][]string{{"a"}, {"b"}, {"c"}}
		assert.Equal(t, expected, result)
	})

	t.Run("CombinationsChooseAll", func(t *testing.T) {
		slice := []int{1, 2}
		result := Combinations(slice, 2)

		expected := [][]int{{1, 2}}
		assert.Equal(t, expected, result)
	})

	t.Run("CombinationsInvalidR", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Combinations(slice, 0)

		assert.Empty(t, result)

		result = Combinations(slice, 4)
		assert.Empty(t, result)
	})
}

package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	t.Run("TransposeNormal", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}
		result := Transpose(matrix)

		expected := [][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("TransposeSquare", func(t *testing.T) {
		matrix := [][]int{
			{1, 2},
			{3, 4},
		}
		result := Transpose(matrix)

		expected := [][]int{
			{1, 3},
			{2, 4},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("TransposeEmpty", func(t *testing.T) {
		var empty [][]int
		result := Transpose(empty)

		assert.Empty(t, result)
	})

	t.Run("TransposeSingleRow", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3, 4}}
		result := Transpose(matrix)

		expected := [][]int{{1}, {2}, {3}, {4}}
		assert.Equal(t, expected, result)
	})

	t.Run("TransposeSingleColumn", func(t *testing.T) {
		matrix := [][]int{{1}, {2}, {3}}
		result := Transpose(matrix)

		expected := [][]int{{1, 2, 3}}
		assert.Equal(t, expected, result)
	})
}

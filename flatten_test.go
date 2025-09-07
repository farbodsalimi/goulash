package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	t.Run("FlattenNumbers", func(t *testing.T) {
		nested := [][]int{{1, 2}, {3, 4}, {5}}
		result := Flatten(nested)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("FlattenStrings", func(t *testing.T) {
		nested := [][]string{{"hello", "world"}, {"go"}, {"programming"}}
		result := Flatten(nested)

		assert.Equal(t, []string{"hello", "world", "go", "programming"}, result)
	})

	t.Run("FlattenEmpty", func(t *testing.T) {
		var empty [][]int
		result := Flatten(empty)

		assert.Empty(t, result)
	})

	t.Run("FlattenWithEmptySlices", func(t *testing.T) {
		nested := [][]int{{1, 2}, {}, {3}}
		result := Flatten(nested)

		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

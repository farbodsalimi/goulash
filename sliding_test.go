package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliding(t *testing.T) {
	t.Run("SlidingNormal", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sliding(numbers, 3)

		expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
		assert.Equal(t, expected, result)
	})

	t.Run("SlidingSizeTwo", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := Sliding(numbers, 2)

		expected := [][]int{{1, 2}, {2, 3}, {3, 4}}
		assert.Equal(t, expected, result)
	})

	t.Run("SlidingSizeEqualLength", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sliding(numbers, 3)

		expected := [][]int{{1, 2, 3}}
		assert.Equal(t, expected, result)
	})

	t.Run("SlidingSizeGreaterThanLength", func(t *testing.T) {
		numbers := []int{1, 2}
		result := Sliding(numbers, 5)

		expected := [][]int{{1, 2}}
		assert.Equal(t, expected, result)
	})

	t.Run("SlidingZeroSize", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sliding(numbers, 0)

		assert.Empty(t, result)
	})

	t.Run("SlidingEmpty", func(t *testing.T) {
		var empty []int
		result := Sliding(empty, 2)

		assert.Empty(t, result)
	})
}

package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	t.Run("ScanSum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := Scan(numbers, func(acc, n int) int {
			return acc + n
		}, 0)

		expected := []int{0, 1, 3, 6, 10}
		assert.Equal(t, expected, result)
	})

	t.Run("ScanProduct", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		result := Scan(numbers, func(acc, n int) int {
			return acc * n
		}, 1)

		expected := []int{1, 2, 6, 24}
		assert.Equal(t, expected, result)
	})

	t.Run("ScanEmpty", func(t *testing.T) {
		var empty []int
		result := Scan(empty, func(acc, n int) int {
			return acc + n
		}, 10)

		expected := []int{10}
		assert.Equal(t, expected, result)
	})

	t.Run("ScanStrings", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		result := Scan(words, func(acc, s string) string {
			return acc + s
		}, "")

		expected := []string{"", "a", "ab", "abc"}
		assert.Equal(t, expected, result)
	})
}

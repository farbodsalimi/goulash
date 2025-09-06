package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("ProductIntegers", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		result := Product(numbers)

		assert.Equal(t, 24, result)
	})

	t.Run("ProductFloats", func(t *testing.T) {
		numbers := []float64{1.5, 2.0, 3.0}
		result := Product(numbers)

		assert.Equal(t, 9.0, result)
	})

	t.Run("ProductEmpty", func(t *testing.T) {
		var empty []int
		result := Product(empty)

		assert.Equal(t, 0, result)
	})

	t.Run("ProductWithZero", func(t *testing.T) {
		numbers := []int{1, 2, 0, 4}
		result := Product(numbers)

		assert.Equal(t, 0, result)
	})

	t.Run("ProductSingle", func(t *testing.T) {
		numbers := []int{5}
		result := Product(numbers)

		assert.Equal(t, 5, result)
	})
}

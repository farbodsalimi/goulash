package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Run("RangeNormal", func(t *testing.T) {
		result := Range(1, 5)

		assert.Equal(t, []int{1, 2, 3, 4}, result)
	})

	t.Run("RangeFromZero", func(t *testing.T) {
		result := Range(0, 3)

		assert.Equal(t, []int{0, 1, 2}, result)
	})

	t.Run("RangeEmpty", func(t *testing.T) {
		result := Range(5, 5)

		assert.Empty(t, result)
	})

	t.Run("RangeBackwards", func(t *testing.T) {
		result := Range(5, 1)

		assert.Empty(t, result)
	})

	t.Run("RangeUint", func(t *testing.T) {
		result := Range(uint(2), uint(6))

		assert.Equal(t, []uint{2, 3, 4, 5}, result)
	})
}

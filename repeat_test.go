package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("RepeatNumber", func(t *testing.T) {
		result := Repeat(42, 3)

		assert.Equal(t, []int{42, 42, 42}, result)
	})

	t.Run("RepeatString", func(t *testing.T) {
		result := Repeat("hello", 2)

		assert.Equal(t, []string{"hello", "hello"}, result)
	})

	t.Run("RepeatZero", func(t *testing.T) {
		result := Repeat("test", 0)

		assert.Empty(t, result)
	})

	t.Run("RepeatNegative", func(t *testing.T) {
		result := Repeat(123, -1)

		assert.Empty(t, result)
	})

	t.Run("RepeatOne", func(t *testing.T) {
		result := Repeat("once", 1)

		assert.Equal(t, []string{"once"}, result)
	})
}

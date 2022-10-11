package goulash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	t.Parallel()

	actual := Fold[int]([]int{1, 2, 3}, 0, func(a int, b int) int {
		return a + b
	})
	assert.Equal(t, 6, actual)

	actual2 := Fold[string]([]string{"a", "b", "c"}, "", func(a string, b string) string {
		if a == "" {
			return fmt.Sprintf("%s", b)
		}
		return fmt.Sprintf("%s:%s", a, b)
	})
	assert.Equal(t, "a:b:c", actual2)
}

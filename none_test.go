package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoneInt(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "NoneMatching",
			input:     []int{1, 3, 5},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  true,
		},
		{
			name:      "SomeMatching",
			input:     []int{1, 2, 3},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  false,
		},
		{
			name:      "NoneEmpty",
			input:     []int{},
			predicate: func(n int) bool { return n > 0 },
			expected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, None(tc.input, tc.predicate))
		})
	}
}

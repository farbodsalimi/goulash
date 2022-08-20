package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []int
		fn     func(int) bool
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			expect: []int{2, 4, 6, 8},
			fn: func(n int) bool {
				return n%2 == 0
			},
		},
		{
			input:  []int{2, 4, 6, 8},
			expect: nil,
			fn: func(n int) bool {
				return n%2 == 1
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Filter(testCase.input, testCase.fn))
		})
	}
}

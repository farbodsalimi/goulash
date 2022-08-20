package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "Empty slices",
			input:  []int{},
			expect: 0,
		},
		{
			name:   "Slice with odd number of elements",
			input:  []int{3, 4, 5, 1, 2},
			expect: 1,
		},
		{
			name:   "Slice with even number of elements",
			input:  []int{3, 0, 1, 2},
			expect: 0,
		},
		{
			name:   "Slice with all negative numbers",
			input:  []int{-3, -1, -4, -2},
			expect: -4,
		},
		{
			name:   "Slice with one negative number",
			input:  []int{2, -1, 0, 1},
			expect: -1,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Min(testCase.input...))
		})
	}
}

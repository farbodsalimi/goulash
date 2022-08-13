package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{},
			expect: 0,
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			expect: 1,
		},
		{
			input:  []int{3, 0, 1, 2},
			expect: 0,
		},
		{
			input:  []int{-3, -1, -4, -2},
			expect: -4,
		},
		{
			input:  []int{2, -1, 0, 1},
			expect: -1,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Min(testCase.input...))
	}
}

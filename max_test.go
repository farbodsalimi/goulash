package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
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
			expect: 5,
		},
		{
			input:  []int{3, 0, 1, 2},
			expect: 3,
		},
		{
			input:  []int{-3, -1, -4, -2},
			expect: -1,
		},
		{
			input:  []int{2, -1, 0, 1},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Max(testCase.input...))
	}
}

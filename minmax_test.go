package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {
	testCases := []struct {
		input     []int
		expectMin int
		expectMax int
	}{
		{
			input:     []int{},
			expectMin: 0,
			expectMax: 0,
		},
		{
			input:     []int{3, 4, 5, 1, 2},
			expectMin: 1,
			expectMax: 5,
		},
		{
			input:     []int{3, 0, 1, 2},
			expectMin: 0,
			expectMax: 3,
		},
		{
			input:     []int{-3, -1, -4, -2},
			expectMin: -4,
			expectMax: -1,
		},
		{
			input:     []int{2, -1, 0, 1},
			expectMin: -1,
			expectMax: 2,
		},
	}

	for _, testCase := range testCases {
		actualMin, actualMax := MinMax(testCase.input...)
		assert.Equal(t, testCase.expectMin, actualMin)
		assert.Equal(t, testCase.expectMax, actualMax)
	}
}

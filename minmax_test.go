package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		expectMin int
		expectMax int
	}{
		{
			name:      "Empty slice",
			input:     []int{},
			expectMin: 0,
			expectMax: 0,
		},
		{
			name:      "Slice with odd number of elements",
			input:     []int{3, 4, 5, 1, 2},
			expectMin: 1,
			expectMax: 5,
		},
		{
			name:      "Slice with even number of elements",
			input:     []int{3, 0, 1, 2},
			expectMin: 0,
			expectMax: 3,
		},
		{
			name:      "Slice with all negative numbers",
			input:     []int{-3, -1, -4, -2},
			expectMin: -4,
			expectMax: -1,
		},
		{
			name:      "Slice with one negative number",
			input:     []int{2, -1, 0, 1},
			expectMin: -1,
			expectMax: 2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualMin, actualMax := MinMax(testCase.input...)
			assert.Equal(t, testCase.expectMin, actualMin)
			assert.Equal(t, testCase.expectMax, actualMax)
		})
	}
}

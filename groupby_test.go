package goulash

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	testCases := []struct {
		input  []float64
		expect map[float64][]float64
	}{
		{
			input: []float64{6.1, 4.2, 6.3},
			expect: map[float64][]float64{
				4: {4.2},
				6: {6.1, 6.3},
			},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, GroupBy(testCase.input, math.Floor))
	}
}

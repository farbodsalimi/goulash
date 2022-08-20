package goulash

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []float64
	}{
		{
			input:  []int{1, 2, 3},
			expect: []float64{10, 100, 1000},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Map(testCase.input, math.Pow10))
		})
	}
}

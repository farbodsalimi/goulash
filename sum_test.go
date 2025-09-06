package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			name: "SumIntegers",
			testFunc: func(t *testing.T) {
				numbers := []int{1, 2, 3, 4, 5}
				result := Sum(numbers)
				assert.Equal(t, 15, result)
			},
		},
		{
			name: "SumFloats",
			testFunc: func(t *testing.T) {
				numbers := []float64{1.5, 2.5, 3.0}
				result := Sum(numbers)
				assert.Equal(t, 7.0, result)
			},
		},
		{
			name: "SumEmpty",
			testFunc: func(t *testing.T) {
				var empty []int
				result := Sum(empty)
				assert.Equal(t, 0, result)
			},
		},
		{
			name: "SumNegative",
			testFunc: func(t *testing.T) {
				numbers := []int{-1, -2, -3}
				result := Sum(numbers)
				assert.Equal(t, -6, result)
			},
		},
		{
			name: "SumStrings",
			testFunc: func(t *testing.T) {
				strings := []string{"hello", " ", "world"}
				result := Sum(strings)
				assert.Equal(t, "hello world", result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testFunc)
	}
}

package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	testCases := []struct {
		input  []int
		expect [][]int
	}{
		{
			input:  []int{1, 2, 3},
			expect: [][]int{{0, 1}, {1, 2}, {2, 3}},
		},
	}

	for _, testCase := range testCases {
		var actual [][]int
		ForEach(testCase.input, func(value int, args ...any) {
			index := args[0].(int)
			actual = append(actual, []int{index, value})
		})

		assert.Equal(t, testCase.expect, actual)
	}
}

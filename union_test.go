package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	testCases := []struct {
		input  [][]int
		expect []int
	}{
		{
			input:  [][]int{{1, 2, 3}, {1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			input:  [][]int{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}},
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  [][]int{{1, 1, 1, 1}, {1}},
			expect: []int{1},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Union(testCase.input...))
	}
}

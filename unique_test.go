package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{1, 1, 1, 1, 2, 3},
			expect: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Unique(testCase.input))
	}
}

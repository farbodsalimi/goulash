package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContact(t *testing.T) {
	testCases := []struct {
		input1 []int
		input2 []int
		expect []int
	}{
		{
			input1: []int{1, 2, 3},
			input2: []int{4, 5, 6},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, Concat(testCase.input1, testCase.input2), testCase.expect)
	}
}

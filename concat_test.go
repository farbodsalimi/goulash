package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
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
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				testCase.expect,
				Concat(testCase.input1, testCase.input2),
			)
		})
	}
}

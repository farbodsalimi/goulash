package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []int
		input2 []string
		expect []Pair[int, string]
	}{
		{
			name:   "Empty slices",
			input1: []int{1, 2, 3},
			input2: []string{"a", "b", "c"},
			expect: []Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
				{First: 3, Second: "c"},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Zip(testCase.input1, testCase.input2))
		})
	}
}

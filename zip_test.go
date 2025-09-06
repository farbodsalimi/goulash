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
			name:   "Same length slices",
			input1: []int{1, 2, 3},
			input2: []string{"a", "b", "c"},
			expect: []Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
				{First: 3, Second: "c"},
			},
		},
		{
			name:   "Empty slices",
			input1: []int{},
			input2: []string{},
			expect: []Pair[int, string]{},
		},
		{
			name:   "Different lengths - input1 shorter",
			input1: []int{1, 2},
			input2: []string{"a", "b", "c"},
			expect: []Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
			},
		},
		{
			name:   "Different lengths - input2 shorter",
			input1: []int{1, 2, 3},
			input2: []string{"a", "b"},
			expect: []Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
			},
		},
		{
			name:   "Nil slices",
			input1: nil,
			input2: nil,
			expect: []Pair[int, string]{},
		},
		{
			name:   "Nil input1",
			input1: nil,
			input2: []string{"a", "b", "c"},
			expect: []Pair[int, string]{},
		},
		{
			name:   "Nil input2",
			input1: []int{1, 2, 3},
			input2: nil,
			expect: []Pair[int, string]{},
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

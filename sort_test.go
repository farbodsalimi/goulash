package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "Empty slice",
			input:  []int{},
			expect: []int{},
		},
		{
			name:   "Slice with 1 element",
			input:  []int{1},
			expect: []int{1},
		},
		{
			name:   "Slices with even number of elements",
			input:  []int{3, 0, 1, 2},
			expect: []int{0, 1, 2, 3},
		},
		{
			name:   "Slices with all negative numbers",
			input:  []int{-3, -1, -4, -2},
			expect: []int{-4, -3, -2, -1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Sort(testCase.input))
		})
	}
}

func TestSortString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "Empty slice",
			input:  []string{},
			expect: []string{},
		},
		{
			name:   "Slice with 1 element",
			input:  []string{"a"},
			expect: []string{"a"},
		},
		{
			name:   "Slices with even number of elements",
			input:  []string{"d", "a", "b", "c"},
			expect: []string{"a", "b", "c", "d"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Sort(testCase.input))
		})
	}
}

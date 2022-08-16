package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortInt(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{},
			expect: []int{},
		},
		{
			input:  []int{1},
			expect: []int{1},
		},
		{
			input:  []int{3, 0, 1, 2},
			expect: []int{0, 1, 2, 3},
		},
		{
			input:  []int{-3, -1, -4, -2},
			expect: []int{-4, -3, -2, -1},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Sort(testCase.input))
	}
}

func TestSortString(t *testing.T) {
	testCases := []struct {
		input  []string
		expect []string
	}{
		{
			input:  []string{},
			expect: []string{},
		},
		{
			input:  []string{"a"},
			expect: []string{"a"},
		},
		{
			input:  []string{"d", "a", "b", "c"},
			expect: []string{"a", "b", "c", "d"},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Sort(testCase.input))
	}
}
package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  [][]int
		expect []int
	}{
		{
			name:   "Empty slices",
			input:  [][]int{},
			expect: []int{},
		},
		{
			name:   "Slices with identical elements",
			input:  [][]int{{1, 2, 3}, {1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "Slices with odd number of elements",
			input:  [][]int{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}},
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "Slices with even number of elements",
			input:  [][]int{{1, 1, 1, 1}, {1}},
			expect: []int{1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Union(testCase.input...))
		})
	}
}

func TestUnionUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  [][]uint
		expect []uint
	}{
		{
			name:   "Empty slices",
			input:  [][]uint{},
			expect: []uint{},
		},
		{
			name:   "Slices with identical elements",
			input:  [][]uint{{1, 2, 3}, {1, 2, 3}},
			expect: []uint{1, 2, 3},
		},
		{
			name:   "Slices with odd number of elements",
			input:  [][]uint{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}},
			expect: []uint{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "Slices with even number of elements",
			input:  [][]uint{{1, 1, 1, 1}, {1}},
			expect: []uint{1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Union(testCase.input...))
		})
	}
}

func TestUnionFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  [][]float32
		expect []float32
	}{
		{
			name:   "Empty slices",
			input:  [][]float32{},
			expect: []float32{},
		},
		{
			name:   "Slices with identical elements",
			input:  [][]float32{{1, 2, 3}, {1, 2, 3}},
			expect: []float32{1, 2, 3},
		},
		{
			name:   "Slices with odd number of elements",
			input:  [][]float32{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}},
			expect: []float32{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "Slices with even number of elements",
			input:  [][]float32{{1, 1, 1, 1}, {1}},
			expect: []float32{1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Union(testCase.input...))
		})
	}
}

func TestUnionFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  [][]float64
		expect []float64
	}{
		{
			name:   "Empty slices",
			input:  [][]float64{},
			expect: []float64{},
		},
		{
			name:   "Slices with identical elements",
			input:  [][]float64{{1, 2, 3}, {1, 2, 3}},
			expect: []float64{1, 2, 3},
		},
		{
			name:   "Slices with odd number of elements",
			input:  [][]float64{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}},
			expect: []float64{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "Slices with even number of elements",
			input:  [][]float64{{1, 1, 1, 1}, {1}},
			expect: []float64{1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Union(testCase.input...))
		})
	}
}

func TestUnionString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  [][]string
		expect []string
	}{
		{
			name:   "Empty slices",
			input:  [][]string{},
			expect: []string{},
		},
		{
			name:   "Slices with identical elements",
			input:  [][]string{{"a", "b", "c"}, {"a", "b", "c"}},
			expect: []string{"a", "b", "c"},
		},
		{
			name:   "Slices with odd number of elements",
			input:  [][]string{{"a", "b"}, {"b", "c", "d"}, {"c", "d", "e", "f", "g"}},
			expect: []string{"a", "b", "c", "d", "e", "f", "g"},
		},
		{
			name:   "Slices with even number of elements",
			input:  [][]string{{"a", "a", "a", "a"}, {"a"}},
			expect: []string{"a"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Union(testCase.input...))
		})
	}
}

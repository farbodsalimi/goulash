package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueInt(t *testing.T) {
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
			name:   "Slices with odd number of elements",
			input:  []int{1, 1, 1, 1, 2, 3},
			expect: []int{1, 2, 3},
		},
		{
			name:   "Slices with even number of elements",
			input:  []int{1, 1, 2, 2, 3, 3, 4, 4},
			expect: []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Unique(testCase.input))
		})
	}
}

func TestUniqueUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint
		expect []uint
	}{
		{
			name:   "Empty slice",
			input:  []uint{},
			expect: []uint{},
		},
		{
			name:   "Slice with 1 element",
			input:  []uint{1},
			expect: []uint{1},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []uint{1, 1, 1, 1, 2, 3},
			expect: []uint{1, 2, 3},
		},
		{
			name:   "Slices with even number of elements",
			input:  []uint{1, 1, 2, 2, 3, 3, 4, 4},
			expect: []uint{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Unique(testCase.input))
		})
	}
}

func TestUniqueFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect []float32
	}{
		{
			name:   "Empty slice",
			input:  []float32{},
			expect: []float32{},
		},
		{
			name:   "Slice with 1 element",
			input:  []float32{1},
			expect: []float32{1},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float32{1, 1, 1, 1, 2, 3},
			expect: []float32{1, 2, 3},
		},
		{
			name:   "Slices with even number of elements",
			input:  []float32{1, 1, 2, 2, 3, 3, 4, 4},
			expect: []float32{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Unique(testCase.input))
		})
	}
}

func TestUniqueFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect []float64
	}{
		{
			name:   "Empty slice",
			input:  []float64{},
			expect: []float64{},
		},
		{
			name:   "Slice with 1 element",
			input:  []float64{1},
			expect: []float64{1},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float64{1, 1, 1, 1, 2, 3},
			expect: []float64{1, 2, 3},
		},
		{
			name:   "Slices with even number of elements",
			input:  []float64{1, 1, 2, 2, 3, 3, 4, 4},
			expect: []float64{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Unique(testCase.input))
		})
	}
}

func TestUniqueString(t *testing.T) {
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
			name:   "Slices with odd number of elements",
			input:  []string{"a", "a", "a", "a", "b", "c"},
			expect: []string{"a", "b", "c"},
		},
		{
			name:   "Slices with even number of elements",
			input:  []string{"a", "a", "b", "b", "c", "c", "d", "d"},
			expect: []string{"a", "b", "c", "d"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Unique(testCase.input))
		})
	}
}

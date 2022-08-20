package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "Empty slices",
			input:  []int{},
			expect: 0,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []int{3, 4, 5, 1, 2},
			expect: 5,
		},
		{
			name:   "Slices with even number of elements",
			input:  []int{3, 0, 1, 2},
			expect: 3,
		},
		{
			name:   "Slices with all negative numbers",
			input:  []int{-3, -1, -4, -2},
			expect: -1,
		},
		{
			name:   "Slices with negative numbers and zero",
			input:  []int{2, -1, 0, 1},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Max(testCase.input...))
		})
	}
}

func TestMaxUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint
		expect uint
	}{
		{
			name:   "Empty slices",
			input:  []uint{},
			expect: 0,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []uint{3, 4, 5, 1, 2},
			expect: 5,
		},
		{
			name:   "Slices with even number of elements",
			input:  []uint{3, 0, 1, 2},
			expect: 3,
		},
		{
			name:   "Slices with zero elements",
			input:  []uint{2, 0, 1},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Max(testCase.input...))
		})
	}
}

func TestMaxFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect float32
	}{
		{
			name:   "Empty slices",
			input:  []float32{},
			expect: 0,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float32{3, 4, 5, 1, 2},
			expect: 5,
		},
		{
			name:   "Slices with even number of elements",
			input:  []float32{3, 0, 1, 2},
			expect: 3,
		},
		{
			name:   "Slices with all negative numbers",
			input:  []float32{-3, -1, -4, -2},
			expect: -1,
		},
		{
			name:   "Slices with negative numbers and zero",
			input:  []float32{2, -1, 0, 1},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Max(testCase.input...))
		})
	}
}

func TestMaxFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect float64
	}{
		{
			name:   "Empty slices",
			input:  []float64{},
			expect: 0,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float64{3, 4, 5, 1, 2},
			expect: 5,
		},
		{
			name:   "Slices with even number of elements",
			input:  []float64{3, 0, 1, 2},
			expect: 3,
		},
		{
			name:   "Slices with all negative numbers",
			input:  []float64{-3, -1, -4, -2},
			expect: -1,
		},
		{
			name:   "Slices with negative numbers and zero",
			input:  []float64{2, -1, 0, 1},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Max(testCase.input...))
		})
	}
}

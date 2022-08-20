package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompactInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "Empty slices",
			input:  []int{},
			expect: []int{},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []int{3, 4, 5, 1, 2},
			expect: []int{3, 4, 5, 1, 2},
		},
		{
			name:   "Slices with even number of elements",
			input:  []int{3, 0, 1, 2},
			expect: []int{3, 1, 2},
		},
		{
			name:   "Slices with negative numbers",
			input:  []int{0, -1, 0, -2},
			expect: []int{-1, -2},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Compact(testCase.input))
		})
	}
}

func TestCompactUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint
		expect []uint
	}{
		{
			name:   "Empty slices",
			input:  []uint{},
			expect: []uint{},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []uint{3, 4, 5, 1, 2},
			expect: []uint{3, 4, 5, 1, 2},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []uint{0, 1, 2},
			expect: []uint{1, 2},
		},
		{
			name:   "Slices with even number of elements",
			input:  []uint{0, 1, 0, 2},
			expect: []uint{1, 2},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Compact(testCase.input))
		})
	}
}

func TestCompactFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect []float32
	}{
		{
			name:   "Empty slices",
			input:  []float32{},
			expect: []float32{},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: []float32{3.1, 4.1, 5.1, 1.1, 2.1},
		},
		{
			name:   "Slices with even number of elements",
			input:  []float32{0.1, 0.0, 0, 2.1},
			expect: []float32{0.1, 2.1},
		},
		{
			name:   "Slices with negative numbers",
			input:  []float32{0, -1.111, -0.0, -2.222},
			expect: []float32{-1.111, -2.222},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Compact(testCase.input))
		})
	}
}

func TestCompactFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect []float64
	}{
		{
			name:   "Empty slices",
			input:  []float64{},
			expect: []float64{},
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: []float64{3.1, 4.1, 5.1, 1.1, 2.1},
		},
		{
			name:   "Slices with even number of elements",
			input:  []float64{0.1, 0.0, 0, 2.1},
			expect: []float64{0.1, 2.1},
		},
		{
			name:   "Slices with negative numbers",
			input:  []float64{0, -1.111, -0.0, -2.222},
			expect: []float64{-1.111, -2.222},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Compact(testCase.input))
		})
	}
}

func TestCompactString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "Empty slices",
			input:  []string{},
			expect: []string{},
		},
		{
			name:   "Slices with even number of elements",
			input:  []string{"a", "b", "c", "d"},
			expect: []string{"a", "b", "c", "d"},
		},
		{
			name:   "Slice with an element with one space character",
			input:  []string{"", "", string(""), " "},
			expect: []string{" "},
		},
		{
			name:   "Two empty string elements",
			input:  []string{"a", "", "c", ""},
			expect: []string{"a", "c"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Compact(testCase.input))
		})
	}
}

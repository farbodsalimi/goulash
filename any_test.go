package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{
			name:   "Empty slices",
			input:  []int{},
			expect: false,
		},
		{
			name:   "Slice with all zero elements",
			input:  []int{0, 0, 0},
			expect: false,
		},
		{
			name:   "Slice with no zero elements",
			input:  []int{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			name:   "Slice with one zero",
			input:  []int{3, 0, 1, 2},
			expect: true,
		},
		{
			name:   "Slice with one zero and negative numbers",
			input:  []int{0, -1, 0, -2},
			expect: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Any(testCase.input))
		})
	}
}

func TestAnyUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint
		expect bool
	}{
		{
			name:   "Empty slices",
			input:  []uint{},
			expect: false,
		},
		{
			name:   "Slice with all zero elements",
			input:  []uint{0, 0, 0},
			expect: false,
		},
		{
			name:   "Slice with no zero elements",
			input:  []uint{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			name:   "Slice with one zero",
			input:  []uint{3, 0, 1, 2},
			expect: true,
		},
		{
			name:   "Slice with multiple zero elements",
			input:  []uint{0, 1, 0, 2},
			expect: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Any(testCase.input))
		})
	}
}

func TestAnyFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect bool
	}{
		{
			name:   "Empty slices",
			input:  []float32{},
			expect: false,
		},
		{
			name:   "Slice with all zero elements",
			input:  []float32{0, 0, 0},
			expect: false,
		},
		{
			name:   "Slice with no zero elements",
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			name:   "Slice with float zero",
			input:  []float32{0.1, 0.0, 0, 2.1},
			expect: true,
		},
		{
			name:   "Slice with float negative zero",
			input:  []float32{0, -1.111, -0.0, -2.222},
			expect: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Any(testCase.input))
		})
	}
}

func TestAnyFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect bool
	}{
		{
			name:   "Empty slices",
			input:  []float64{},
			expect: false,
		},
		{
			name:   "Slice with all zero elements",
			input:  []float64{0, 0, 0},
			expect: false,
		},
		{
			name:   "Slice with no zero elements",
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			name:   "Slice with float zero",
			input:  []float64{0.1, 0.0, 0, 2.1},
			expect: true,
		},
		{
			name:   "Slice with float negative zero",
			input:  []float64{0, -1.111, -0.0, -2.222},
			expect: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Any(testCase.input))
		})
	}
}

func TestAnyString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "Empty slices",
			input:  []string{},
			expect: false,
		},
		{
			name:   "Slice with all empty string elements",
			input:  []string{"", "", ""},
			expect: false,
		},
		{
			name:   "Slice with an element with one space character",
			input:  []string{"", "", string(""), " "},
			expect: true,
		},
		{
			name:   "Slice with no empty string elements",
			input:  []string{"a", "b", "c", "d"},
			expect: true,
		},
		{
			name:   "Slice with multiple empty string elements",
			input:  []string{"a", "", "c", ""},
			expect: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Any(testCase.input))
		})
	}
}

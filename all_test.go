package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{
			name:   "Empty slice",
			input:  []int{},
			expect: true,
		},
		{
			name:   "Valid slice",
			input:  []int{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			name:   "One zero element",
			input:  []int{3, 0, 1, 2},
			expect: false,
		},
		{
			name:   "Two zero elements",
			input:  []int{0, -1, 0, -2},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, All(testCase.input))
		})
	}
}

func TestAllUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint
		expect bool
	}{
		{
			name:   "Empty slice",
			input:  []uint{},
			expect: true,
		},
		{
			name:   "Valid slice",
			input:  []uint{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			name:   "One zero element",
			input:  []uint{3, 0, 1, 2},
			expect: false,
		},
		{
			name:   "Two zero elements",
			input:  []uint{0, 1, 0, 2},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, All(testCase.input))
		})
	}
}

func TestAllFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect bool
	}{
		{
			name:   "Empty slice",
			input:  []float32{},
			expect: true,
		},
		{
			name:   "Valid slice",
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			name:   "Two zero elements",
			input:  []float32{0.1, 0.0, 0, 2.1},
			expect: false,
		},
		{
			name:   "One negative zero element",
			input:  []float32{-1.111, -0.0, -2.222},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, All(testCase.input))
		})
	}
}

func TestAllFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect bool
	}{
		{
			name:   "Empty slice",
			input:  []float64{},
			expect: true,
		},
		{
			name:   "Valid slice",
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			name:   "Two zero elements",
			input:  []float64{0.1, 0.0, 0, 2.1},
			expect: false,
		},
		{
			name:   "One negative zero element",
			input:  []float64{-1.111, -0.0, -2.222},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, All(testCase.input))
		})
	}
}

func TestAllString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "Empty slice",
			input:  []string{},
			expect: true,
		},
		{
			name:   "Valid slice",
			input:  []string{"a", "b", "c", "d"},
			expect: true,
		},
		{
			name:   "All elements are empty string",
			input:  []string{"", "", string(""), " "},
			expect: false,
		},
		{
			name:   "Two empty string elements",
			input:  []string{"a", "", "c", ""},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, All(testCase.input))
		})
	}
}

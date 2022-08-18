package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllInt(t *testing.T) {
	testCases := []struct {
		input  []int
		expect bool
	}{
		{
			input:  []int{},
			expect: true,
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			input:  []int{3, 0, 1, 2},
			expect: false,
		},
		{
			input:  []int{0, -1, 0, -2},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, All(testCase.input))
	}
}

func TestAllUint(t *testing.T) {
	testCases := []struct {
		input  []uint
		expect bool
	}{
		{
			input:  []uint{},
			expect: true,
		},
		{
			input:  []uint{3, 4, 5, 1, 2},
			expect: true,
		},
		{
			input:  []uint{3, 0, 1, 2},
			expect: false,
		},
		{
			input:  []uint{0, 1, 0, 2},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, All(testCase.input))
	}
}

func TestAllFloat32(t *testing.T) {
	testCases := []struct {
		input  []float32
		expect bool
	}{
		{
			input:  []float32{},
			expect: true,
		},
		{
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			input:  []float32{0.1, 0.0, 0, 2.1},
			expect: false,
		},
		{
			input:  []float32{0, -1.111, -0.0, -2.222},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, All(testCase.input))
	}
}

func TestAllFloat64(t *testing.T) {
	testCases := []struct {
		input  []float64
		expect bool
	}{
		{
			input:  []float64{},
			expect: true,
		},
		{
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: true,
		},
		{
			input:  []float64{0.1, 0.0, 0, 2.1},
			expect: false,
		},
		{
			input:  []float64{0, -1.111, -0.0, -2.222},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, All(testCase.input))
	}
}

func TestAllString(t *testing.T) {
	testCases := []struct {
		input  []string
		expect bool
	}{
		{
			input:  []string{},
			expect: true,
		},
		{
			input:  []string{"a", "b", "c", "d"},
			expect: true,
		},
		{
			input:  []string{"", "", string(""), " "},
			expect: false,
		},
		{
			input:  []string{"a", "", "c", ""},
			expect: false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, All(testCase.input))
	}
}

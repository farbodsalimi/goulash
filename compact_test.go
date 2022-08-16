package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompactInt(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{},
			expect: []int{},
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			expect: []int{3, 4, 5, 1, 2},
		},
		{
			input:  []int{3, 0, 1, 2},
			expect: []int{3, 1, 2},
		},
		{
			input:  []int{0, -1, 0, -2},
			expect: []int{-1, -2},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Compact(testCase.input))
	}
}

func TestCompactUint(t *testing.T) {
	testCases := []struct {
		input  []uint
		expect []uint
	}{
		{
			input:  []uint{},
			expect: []uint{},
		},
		{
			input:  []uint{3, 4, 5, 1, 2},
			expect: []uint{3, 4, 5, 1, 2},
		},
		{
			input:  []uint{3, 0, 1, 2},
			expect: []uint{3, 1, 2},
		},
		{
			input:  []uint{0, 1, 0, 2},
			expect: []uint{1, 2},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Compact(testCase.input))
	}
}

func TestCompactFloat32(t *testing.T) {
	testCases := []struct {
		input  []float32
		expect []float32
	}{
		{
			input:  []float32{},
			expect: []float32{},
		},
		{
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: []float32{3.1, 4.1, 5.1, 1.1, 2.1},
		},
		{
			input:  []float32{0.1, 0.0, 0, 2.1},
			expect: []float32{0.1, 2.1},
		},
		{
			input:  []float32{0, -1.111, -0.0, -2.222},
			expect: []float32{-1.111, -2.222},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Compact(testCase.input))
	}
}

func TestCompactFloat64(t *testing.T) {
	testCases := []struct {
		input  []float64
		expect []float64
	}{
		{
			input:  []float64{},
			expect: []float64{},
		},
		{
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: []float64{3.1, 4.1, 5.1, 1.1, 2.1},
		},
		{
			input:  []float64{0.1, 0.0, 0, 2.1},
			expect: []float64{0.1, 2.1},
		},
		{
			input:  []float64{0, -1.111, -0.0, -2.222},
			expect: []float64{-1.111, -2.222},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Compact(testCase.input))
	}
}

func TestCompactString(t *testing.T) {
	testCases := []struct {
		input  []string
		expect []string
	}{
		{
			input:  []string{},
			expect: []string{},
		},
		{
			input:  []string{"a", "b", "c", "d"},
			expect: []string{"a", "b", "c", "d"},
		},
		{
			input:  []string{"", "", string(""), " "},
			expect: []string{" "},
		},
		{
			input:  []string{"a", "", "c", ""},
			expect: []string{"a", "c"},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Compact(testCase.input))
	}
}

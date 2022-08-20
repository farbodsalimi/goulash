package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		size   int
		input  []int
		expect [][]int
	}{
		{
			size:   3,
			input:  []int{},
			expect: [][]int{},
		},
		{
			size:   0,
			input:  []int{1},
			expect: [][]int{{1}},
		},
		{
			size:   -1,
			input:  []int{1},
			expect: [][]int{{1}},
		},
		{
			size:   2,
			input:  []int{3, 4, 5, 1, 2},
			expect: [][]int{{3, 4}, {5, 1}, {2}},
		},
		{
			size:   3,
			input:  []int{3, 0, 1, 2},
			expect: [][]int{{3, 0, 1}, {2}},
		},
		{
			size:   1,
			input:  []int{0, -1, 0, -2},
			expect: [][]int{{0}, {-1}, {0}, {-2}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Chunk(testCase.input, testCase.size))
		})
	}
}

func TestChunkUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		size   int
		input  []uint
		expect [][]uint
	}{
		{
			size:   3,
			input:  []uint{},
			expect: [][]uint{},
		},
		{
			size:   0,
			input:  []uint{1},
			expect: [][]uint{{1}},
		},
		{
			size:   -1,
			input:  []uint{1},
			expect: [][]uint{{1}},
		},
		{
			size:   2,
			input:  []uint{3, 4, 5, 1, 2},
			expect: [][]uint{{3, 4}, {5, 1}, {2}},
		},
		{
			size:   3,
			input:  []uint{3, 0, 1, 2},
			expect: [][]uint{{3, 0, 1}, {2}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Chunk(testCase.input, testCase.size))
		})
	}
}

func TestChunkFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		size   int
		input  []float32
		expect [][]float32
	}{
		{
			size:   3,
			input:  []float32{},
			expect: [][]float32{},
		},
		{
			size:   0,
			input:  []float32{1.1},
			expect: [][]float32{{1.1}},
		},
		{
			size:   -1,
			input:  []float32{1.1},
			expect: [][]float32{{1.1}},
		},
		{
			size:   2,
			input:  []float32{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: [][]float32{{3.1, 4.1}, {5.1, 1.1}, {2.1}},
		},
		{
			size:   3,
			input:  []float32{3.1, 0.1, 1.1, 2.1},
			expect: [][]float32{{3.1, 0.1, 1.1}, {2.1}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Chunk(testCase.input, testCase.size))
		})
	}
}

func TestChunkFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		size   int
		input  []float64
		expect [][]float64
	}{
		{
			size:   3,
			input:  []float64{},
			expect: [][]float64{},
		},
		{
			size:   0,
			input:  []float64{1.1},
			expect: [][]float64{{1.1}},
		},
		{
			size:   -1,
			input:  []float64{1.1},
			expect: [][]float64{{1.1}},
		},
		{
			size:   2,
			input:  []float64{3.1, 4.1, 5.1, 1.1, 2.1},
			expect: [][]float64{{3.1, 4.1}, {5.1, 1.1}, {2.1}},
		},
		{
			size:   3,
			input:  []float64{3.1, 0.1, 1.1, 2.1},
			expect: [][]float64{{3.1, 0.1, 1.1}, {2.1}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Chunk(testCase.input, testCase.size))
		})
	}
}

func TestChunkString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		size   int
		input  []string
		expect [][]string
	}{
		{
			size:   3,
			input:  []string{},
			expect: [][]string{},
		},
		{
			size:   0,
			input:  []string{"a"},
			expect: [][]string{{"a"}},
		},
		{
			size:   -1,
			input:  []string{"a"},
			expect: [][]string{{"a"}},
		},
		{
			size:   2,
			input:  []string{"a", "b", "c", "d", "e"},
			expect: [][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			size:   3,
			input:  []string{"a", "b", "c", "d"},
			expect: [][]string{{"a", "b", "c"}, {"d"}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Chunk(testCase.input, testCase.size))
		})
	}
}

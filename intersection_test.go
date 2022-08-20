package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		expect []int
	}{
		{
			name:   "Empty slices",
			input1: []int{},
			input2: []int{},
			expect: []int{},
		},
		{
			name:   "First slice empty",
			input1: []int{},
			input2: []int{1},
			expect: []int{},
		},
		{
			name:   "Second slice empty",
			input1: []int{1},
			input2: []int{},
			expect: []int{},
		},
		{
			name:   "Slices with 1 element",
			input1: []int{1},
			input2: []int{1},
			expect: []int{1},
		},
		{
			name:   "Slices with odd number of elements",
			input1: []int{1, 2, 3},
			input2: []int{1},
			expect: []int{1},
		},
		{
			name:   "Slices with even number of elements",
			input1: []int{-1, 0, 1, 3},
			input2: []int{-2, -3, -4, -1},
			expect: []int{-1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Intersection(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestIntersectionUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []uint
		input2 []uint
		expect []uint
	}{
		{
			name:   "Empty slices",
			input1: []uint{},
			input2: []uint{},
			expect: []uint{},
		},
		{
			name:   "First slice empty",
			input1: []uint{},
			input2: []uint{1},
			expect: []uint{},
		},
		{
			name:   "Second slice empty",
			input1: []uint{1},
			input2: []uint{},
			expect: []uint{},
		},
		{
			name:   "Slices with 1 element",
			input1: []uint{1},
			input2: []uint{1},
			expect: []uint{1},
		},
		{
			name:   "Slices with odd number of elements",
			input1: []uint{1, 2, 3},
			input2: []uint{1},
			expect: []uint{1},
		},
		{
			name:   "Slices with even number of elements",
			input1: []uint{0, 1},
			input2: []uint{2, 3, 4, 0},
			expect: []uint{0},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Intersection(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestIntersectionFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []float32
		input2 []float32
		expect []float32
	}{
		{
			name:   "Empty slices",
			input1: []float32{},
			input2: []float32{},
			expect: []float32{},
		},
		{
			name:   "First slice empty",
			input1: []float32{},
			input2: []float32{1},
			expect: []float32{},
		},
		{
			name:   "Second slice empty",
			input1: []float32{1},
			input2: []float32{},
			expect: []float32{},
		},
		{
			name:   "Slices with 1 element",
			input1: []float32{1},
			input2: []float32{1},
			expect: []float32{1},
		},
		{
			name:   "Slices with odd number of elements",
			input1: []float32{1, 2, 3},
			input2: []float32{1, 3},
			expect: []float32{1, 3},
		},
		{
			name:   "Slices with even number of elements",
			input1: []float32{0, 1},
			input2: []float32{2, 3, 4, 0},
			expect: []float32{0},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Intersection(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestIntersectionFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []float64
		input2 []float64
		expect []float64
	}{
		{
			name:   "Empty slices",
			input1: []float64{},
			input2: []float64{},
			expect: []float64{},
		},
		{
			name:   "First slice empty",
			input1: []float64{},
			input2: []float64{1},
			expect: []float64{},
		},
		{
			name:   "Second slice empty",
			input1: []float64{1},
			input2: []float64{},
			expect: []float64{},
		},
		{
			name:   "Slices with 1 element",
			input1: []float64{1},
			input2: []float64{1},
			expect: []float64{1},
		},
		{
			name:   "Slices with odd number of elements",
			input1: []float64{1, 2, 3},
			input2: []float64{1},
			expect: []float64{1},
		},
		{
			name:   "Slices with even number of elements",
			input1: []float64{0, 1},
			input2: []float64{2, 3, 4, 0},
			expect: []float64{0},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Intersection(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestIntersectionString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []string
		input2 []string
		expect []string
	}{
		{
			name:   "Empty slices",
			input1: []string{},
			input2: []string{},
			expect: []string{},
		},
		{
			name:   "First slice empty",
			input1: []string{},
			input2: []string{"a"},
			expect: []string{},
		},
		{
			name:   "Second slice empty",
			input1: []string{"a"},
			input2: []string{},
			expect: []string{},
		},
		{
			name:   "Slices with 1 element",
			input1: []string{"a"},
			input2: []string{"a"},
			expect: []string{"a"},
		},
		{
			name:   "Slices with odd number of elements",
			input1: []string{"a", "b", "c"},
			input2: []string{"c"},
			expect: []string{"c"},
		},
		{
			name:   "Slices with even number of elements",
			input1: []string{"a", "b", "c", "d"},
			input2: []string{"c", "d"},
			expect: []string{"c", "d"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Intersection(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

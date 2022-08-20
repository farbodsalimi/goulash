package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		expect []int
	}{
		{
			input1: []int{},
			input2: []int{},
			expect: []int{},
		},
		{
			input1: []int{},
			input2: []int{1},
			expect: []int{1},
		},
		{
			input1: []int{1},
			input2: []int{},
			expect: []int{1},
		},
		{
			input1: []int{1},
			input2: []int{1},
			expect: []int{},
		},
		{
			input1: []int{1, 2, 3},
			input2: []int{1, 3},
			expect: []int{2},
		},
		{
			input1: []int{-1, 0, 1},
			input2: []int{2, 3, 4},
			expect: []int{-1, 0, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Difference(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestDifferenceUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []uint
		input2 []uint
		expect []uint
	}{
		{
			input1: []uint{},
			input2: []uint{},
			expect: []uint{},
		},
		{
			input1: []uint{},
			input2: []uint{1},
			expect: []uint{1},
		},
		{
			input1: []uint{1},
			input2: []uint{},
			expect: []uint{1},
		},
		{
			input1: []uint{1},
			input2: []uint{1},
			expect: []uint{},
		},
		{
			input1: []uint{1, 2, 3},
			input2: []uint{1, 3},
			expect: []uint{2},
		},
		{
			input1: []uint{0, 1},
			input2: []uint{2, 3, 4},
			expect: []uint{0, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Difference(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestDifferenceFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []float32
		input2 []float32
		expect []float32
	}{
		{
			input1: []float32{},
			input2: []float32{},
			expect: []float32{},
		},
		{
			input1: []float32{},
			input2: []float32{1},
			expect: []float32{1},
		},
		{
			input1: []float32{1},
			input2: []float32{},
			expect: []float32{1},
		},
		{
			input1: []float32{1},
			input2: []float32{1},
			expect: []float32{},
		},
		{
			input1: []float32{1, 2, 3},
			input2: []float32{1, 3},
			expect: []float32{2},
		},
		{
			input1: []float32{0, 1},
			input2: []float32{2, 3, 4},
			expect: []float32{0, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Difference(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestDifferenceFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []float64
		input2 []float64
		expect []float64
	}{
		{
			input1: []float64{},
			input2: []float64{},
			expect: []float64{},
		},
		{
			input1: []float64{},
			input2: []float64{1},
			expect: []float64{1},
		},
		{
			input1: []float64{1},
			input2: []float64{},
			expect: []float64{1},
		},
		{
			input1: []float64{1},
			input2: []float64{1},
			expect: []float64{},
		},
		{
			input1: []float64{1, 2, 3},
			input2: []float64{1, 3},
			expect: []float64{2},
		},
		{
			input1: []float64{0, 1},
			input2: []float64{2, 3, 4},
			expect: []float64{0, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Difference(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

func TestDifferenceString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input1 []string
		input2 []string
		expect []string
	}{
		{
			input1: []string{},
			input2: []string{},
			expect: []string{},
		},
		{
			input1: []string{"a"},
			input2: []string{"a"},
			expect: []string{},
		},
		{
			input1: []string{"a", "b"},
			input2: []string{"c"},
			expect: []string{"a", "b", "c"},
		},
		{
			input1: []string{"a", "b", "c", "d", "e"},
			input2: []string{"d", "e"},
			expect: []string{"a", "b", "c"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := Difference(testCase.input1, testCase.input2)
			assert.ElementsMatch(t, Sort(testCase.expect), Sort(actual))
		})
	}
}

package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceInt(t *testing.T) {
	t.Parallel()

	var sumIntFn = func(a int, b int) int {
		return a + b
	}

	testCases := []struct {
		name   string
		input  []int
		expect int
		fn     func(a int, b int) int
	}{
		{
			name:   "Empty slices",
			input:  []int{},
			expect: 0,
			fn:     sumIntFn,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []int{1, 2, 3},
			expect: 6,
			fn:     sumIntFn,
		},
		{
			name:   "Slices with even number of elements",
			input:  []int{1, 2, 3, 4},
			expect: 10,
			fn:     sumIntFn,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Reduce(testCase.input, testCase.fn, 0))
		})
	}
}

func TestReduceUint(t *testing.T) {
	t.Parallel()

	var sumUintFn = func(a uint, b uint) uint {
		return a + b
	}

	testCases := []struct {
		name   string
		input  []uint
		expect uint
		fn     func(a uint, b uint) uint
	}{
		{
			name:   "Empty slices",
			input:  []uint{},
			expect: 0,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []uint{1, 2, 3},
			expect: 6,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with even number of elements",
			input:  []uint{1, 2, 3, 4},
			expect: 10,
			fn:     sumUintFn,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Reduce(testCase.input, testCase.fn, 0))
		})
	}
}

func TestReduceFloat32(t *testing.T) {
	t.Parallel()

	var sumUintFn = func(a float32, b float32) float32 {
		return a + b
	}

	testCases := []struct {
		name   string
		input  []float32
		expect float32
		fn     func(a float32, b float32) float32
	}{
		{
			name:   "Empty slices",
			input:  []float32{},
			expect: 0,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float32{1.1, 2.1, 3.1},
			expect: 6.3,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with even number of elements",
			input:  []float32{1.1, 2.1, 3.1, 4.1},
			expect: 10.4,
			fn:     sumUintFn,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Reduce(testCase.input, testCase.fn, 0))
		})
	}
}

func TestReduceFloat64(t *testing.T) {
	t.Parallel()

	var sumUintFn = func(a float64, b float64) float64 {
		return a + b
	}

	testCases := []struct {
		name   string
		input  []float64
		expect float64
		fn     func(a float64, b float64) float64
	}{
		{
			name:   "Empty slices",
			input:  []float64{},
			expect: 0,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with odd number of elements",
			input:  []float64{1.1, 2.1, 3.1},
			expect: 6.300000000000001,
			fn:     sumUintFn,
		},
		{
			name:   "Slices with even number of elements",
			input:  []float64{1.1, 2.1, 3.1, 4.1},
			expect: 10.4,
			fn:     sumUintFn,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Reduce(testCase.input, testCase.fn, 0))
		})
	}
}

func TestReduceString(t *testing.T) {
	t.Parallel()

	var addStringFn = func(a string, b string) string {
		return a + b
	}

	testCases := []struct {
		name      string
		input     []string
		initValue string
		expect    string
		fn        func(a string, b string) string
	}{
		{
			name:      "Empty slices",
			input:     []string{},
			initValue: "",
			expect:    "",
			fn:        addStringFn,
		},
		{
			name:      "Slices with odd number of elements",
			input:     []string{"a", "b", "c"},
			initValue: "letters: ",
			expect:    "letters: abc",
			fn:        addStringFn,
		},
		{
			name:      "Slices with even number of elements",
			input:     []string{"a", "b", "c", "d"},
			initValue: "letters: ",
			expect:    "letters: abcd",
			fn:        addStringFn,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				testCase.expect,
				Reduce(testCase.input, testCase.fn, testCase.initValue),
			)
		})
	}
}

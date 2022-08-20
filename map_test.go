package goulash

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		expect []float64
	}{
		{
			name:   "Empty slice",
			input:  []int{},
			expect: []float64{},
		},
		{
			name:   "Pow10 all elements",
			input:  []int{1, 2, 3},
			expect: []float64{10, 100, 1000},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Map(testCase.input, math.Pow10))
		})
	}
}

func TestMapFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float32
		expect []uint32
	}{
		{
			name:   "Empty slice",
			input:  []float32{},
			expect: []uint32{},
		},
		{
			name:   "Log all elements",
			input:  []float32{1, 2, 3},
			expect: []uint32{1065353216, 1073741824, 1077936128},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Map(testCase.input, math.Float32bits))
		})
	}
}

func TestMapFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []float64
		expect []float64
	}{
		{
			name:   "Empty slice",
			input:  []float64{},
			expect: []float64{},
		},
		{
			name:   "Log all elements",
			input:  []float64{1, 2, 3},
			expect: []float64{0, 0.6931471805599453, 1.0986122886681096},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Map(testCase.input, math.Log))
		})
	}
}

func TestMapString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect [][]string
	}{
		{
			name:   "Empty slice",
			input:  []string{},
			expect: [][]string{},
		},
		{
			name:   "Run Fields on all elements",
			input:  []string{"a", "b", "c"},
			expect: [][]string{{"a"}, {"b"}, {"c"}},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Map(testCase.input, strings.Fields))
		})
	}
}

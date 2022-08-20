package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element int
		input   []int
		expect  bool
	}{
		{
			name:    "Empty slices",
			element: 3,
			input:   []int{},
			expect:  false,
		},
		{
			name:    "Element not found",
			element: 0,
			input:   []int{1},
			expect:  false,
		},
		{
			name:    "Negative lookup element not found",
			element: -1,
			input:   []int{1},
			expect:  false,
		},
		{
			name:    "Positive lookup element found",
			element: 2,
			input:   []int{3, 4, 5, 1, 2},
			expect:  true,
		},
		{
			name:    "Positive lookup element found",
			element: 3,
			input:   []int{3, 0, 1, 2},
			expect:  true,
		},
		{
			name:    "Negative lookup element found",
			element: -1,
			input:   []int{0, -1, 0, -2},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element int
		input   []uint
		expect  bool
	}{
		{
			name:    "Empty slices",
			element: 3,
			input:   []uint{},
			expect:  false,
		},
		{
			name:    "Element not found",
			element: 0,
			input:   []uint{1},
			expect:  false,
		},
		{
			name:    "Negative lookup element not found",
			element: -1,
			input:   []uint{1},
			expect:  false,
		},
		{
			name:    "Positive lookup element found",
			element: 2,
			input:   []uint{3, 4, 5, 1, 2},
			expect:  true,
		},
		{
			name:    "Positive lookup element found",
			element: 3,
			input:   []uint{3, 0, 1, 2},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element float32
		input   []float32
		expect  bool
	}{
		{
			name:    "Empty slices",
			element: 3,
			input:   []float32{},
			expect:  false,
		},
		{
			name:    "Element not found",
			element: 0,
			input:   []float32{1.1},
			expect:  false,
		},
		{
			name:    "Negative lookup element not found",
			element: -1,
			input:   []float32{1.1},
			expect:  false,
		},
		{
			name:    "Negative lookup element found",
			element: -2.1,
			input:   []float32{3.1, 4.1, 5.1, 1.1, -2.1},
			expect:  true,
		},
		{
			name:    "Float lookup element",
			element: 3.1,
			input:   []float32{3.1, 0.1, 1.1, 2.1},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element float64
		input   []float64
		expect  bool
	}{
		{
			name:    "Empty slices",
			element: 3,
			input:   []float64{},
			expect:  false,
		},
		{
			name:    "Element not found",
			element: 0,
			input:   []float64{1.1},
			expect:  false,
		},
		{
			name:    "Negative lookup element not found",
			element: -1,
			input:   []float64{1.1},
			expect:  false,
		},
		{
			name:    "Negative lookup element found",
			element: -2.1,
			input:   []float64{3.1, 4.1, 5.1, 1.1, -2.1},
			expect:  true,
		},
		{
			name:    "Float lookup element",
			element: 3.1,
			input:   []float64{3.1, 0.1, 1.1, 2.1},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsSliceString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element string
		input   []string
		expect  bool
	}{
		{
			name:    "Empty slices",
			element: "a",
			input:   []string{},
			expect:  false,
		},
		{
			name:    "Element not found",
			element: "b",
			input:   []string{"a"},
			expect:  false,
		},
		{
			name:    "Empty lookup element",
			element: "",
			input:   []string{"a"},
			expect:  false,
		},
		{
			name:    "Empty lookup element found",
			element: "",
			input:   []string{""},
			expect:  true,
		},
		{
			name:    "Lookup element found",
			element: "c",
			input:   []string{"a", "b", "c", "d", "e"},
			expect:  true,
		},
		{
			name:    "Lookup element found",
			element: "d",
			input:   []string{"a", "b", "c", "d"},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsMapString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element string
		input   map[string]string
		expect  bool
	}{
		{
			name:    "Empty map",
			element: "key",
			input:   map[string]string{},
			expect:  false,
		},
		{
			name:    "Key not found",
			element: "key1",
			input:   map[string]string{"key": "value"},
			expect:  false,
		},
		{
			name:    "Empty lookup element",
			element: "",
			input:   map[string]string{"key": "value"},
			expect:  false,
		},
		{
			name:    "Map with 1 element",
			element: "key",
			input:   map[string]string{"key": "value"},
			expect:  true,
		},
		{
			name:    "Map with distinct keys/values",
			element: "key1",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
			expect: true,
		},
		{
			name:    "Map with identical values",
			element: "key2",
			input:   map[string]string{"key1": "value", "key2": "value", "key3": "value"},
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

func TestContainsString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		element string
		input   string
		expect  bool
	}{
		{
			name:    "Empty string",
			element: "hello",
			input:   "",
			expect:  false,
		},
		{
			name:    "String not found",
			element: "hello",
			input:   "Hello, World!",
			expect:  false,
		},
		{
			name:    "Partially found",
			element: "llo",
			input:   "Hello, World!",
			expect:  true,
		},
		{
			name:    "String found",
			element: "Hello",
			input:   "Hello, World!",
			expect:  true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Contains(testCase.input, testCase.element))
		})
	}
}

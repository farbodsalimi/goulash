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
			element: 3,
			input:   []int{},
			expect:  false,
		},
		{
			element: 0,
			input:   []int{1},
			expect:  false,
		},
		{
			element: -1,
			input:   []int{1},
			expect:  false,
		},
		{
			element: 2,
			input:   []int{3, 4, 5, 1, 2},
			expect:  true,
		},
		{
			element: 3,
			input:   []int{3, 0, 1, 2},
			expect:  true,
		},
		{
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
			element: 3,
			input:   []uint{},
			expect:  false,
		},
		{
			element: 0,
			input:   []uint{1},
			expect:  false,
		},
		{
			element: -1,
			input:   []uint{1},
			expect:  false,
		},
		{
			element: 2,
			input:   []uint{3, 4, 5, 1, 2},
			expect:  true,
		},
		{
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
			element: 3,
			input:   []float32{},
			expect:  false,
		},
		{
			element: 0,
			input:   []float32{1.1},
			expect:  false,
		},
		{
			element: -1,
			input:   []float32{1.1},
			expect:  false,
		},
		{
			element: -2.1,
			input:   []float32{3.1, 4.1, 5.1, 1.1, -2.1},
			expect:  true,
		},
		{
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
			element: 3,
			input:   []float64{},
			expect:  false,
		},
		{
			element: 0,
			input:   []float64{1.1},
			expect:  false,
		},
		{
			element: -1,
			input:   []float64{1.1},
			expect:  false,
		},
		{
			element: -2.1,
			input:   []float64{3.1, 4.1, 5.1, 1.1, -2.1},
			expect:  true,
		},
		{
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
			element: "a",
			input:   []string{},
			expect:  false,
		},
		{
			element: "b",
			input:   []string{"a"},
			expect:  false,
		},
		{
			element: "",
			input:   []string{"a"},
			expect:  false,
		},
		{
			element: "",
			input:   []string{""},
			expect:  true,
		},
		{
			element: "c",
			input:   []string{"a", "b", "c", "d", "e"},
			expect:  true,
		},
		{
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
			element: "key",
			input:   map[string]string{},
			expect:  false,
		},
		{
			element: "key1",
			input:   map[string]string{"key": "value"},
			expect:  false,
		},
		{
			element: "",
			input:   map[string]string{"key": "value"},
			expect:  false,
		},
		{
			element: "key",
			input:   map[string]string{"key": "value"},
			expect:  true,
		},
		{
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
			element: "hello",
			input:   "",
			expect:  false,
		},
		{
			element: "hello",
			input:   "Hello, World!",
			expect:  false,
		},
		{
			element: "llo",
			input:   "Hello, World!",
			expect:  true,
		},
		{
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

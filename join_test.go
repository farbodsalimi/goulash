package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		delimiter string
		expect    string
	}{
		{
			name:   "Empty slice",
			input:  []int{},
			expect: "",
		},
		{
			name:      "Empty delimiter",
			input:     []int{1, 2, 3},
			delimiter: "",
			expect:    "123",
		},
		{
			name:      "Comma delimiter",
			input:     []int{1, 2, 3, 1, 2, 3},
			delimiter: ",",
			expect:    "1,2,3,1,2,3",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
		})
	}
}

func TestJoinUint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []uint
		delimiter string
		expect    string
	}{
		{
			name:   "Empty slice",
			input:  []uint{},
			expect: "",
		},
		{
			name:      "Empty delimiter",
			input:     []uint{1, 2, 3},
			delimiter: "",
			expect:    "123",
		},
		{
			name:      "Comma delimiter",
			input:     []uint{1, 2, 3, 1, 2, 3},
			delimiter: ",",
			expect:    "1,2,3,1,2,3",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
		})
	}
}

func TestJoinFloat32(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []float32
		delimiter string
		expect    string
	}{
		{
			name:   "Empty slice",
			input:  []float32{},
			expect: "",
		},
		{
			name:      "Empty delimiter",
			input:     []float32{1.1, 2.1, 3.1},
			delimiter: "",
			expect:    "1.12.13.1",
		},
		{
			name:      "Comma delimiter",
			input:     []float32{1.1, 2.1, 3.1, 1.1, 2.1, 3.1},
			delimiter: ",",
			expect:    "1.1,2.1,3.1,1.1,2.1,3.1",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
		})
	}
}

func TestJoinFloat64(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []float64
		delimiter string
		expect    string
	}{
		{
			name:   "Empty slice",
			input:  []float64{},
			expect: "",
		},
		{
			name:      "Empty delimiter",
			input:     []float64{1.1, 2.1, 3.1},
			delimiter: "",
			expect:    "1.12.13.1",
		},
		{
			name:      "Comma delimiter",
			input:     []float64{1.1, 2.1, 3.1, 1.1, 2.1, 3.1},
			delimiter: ",",
			expect:    "1.1,2.1,3.1,1.1,2.1,3.1",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
		})
	}
}

func TestJoinString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []string
		delimiter string
		expect    string
	}{
		{
			name:   "Empty slice",
			input:  []string{},
			expect: "",
		},
		{
			name:      "Empty delimiter",
			input:     []string{"a", "b", "c"},
			delimiter: "",
			expect:    "abc",
		},
		{
			name:      "Comma delimiter",
			input:     []string{"a", "b", "c", "d"},
			delimiter: ",",
			expect:    "a,b,c,d",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
		})
	}
}

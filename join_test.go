package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinInt(t *testing.T) {
	testCases := []struct {
		input     []int
		delimiter string
		expect    string
	}{
		{
			input:  []int{},
			expect: "",
		},
		{
			input:     []int{1, 2, 3},
			delimiter: "",
			expect:    "123",
		},
		{
			input:     []int{1, 2, 3, 1, 2, 3},
			delimiter: ",",
			expect:    "1,2,3,1,2,3",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
	}
}

func TestJoinUint(t *testing.T) {
	testCases := []struct {
		input     []uint
		delimiter string
		expect    string
	}{
		{
			input:  []uint{},
			expect: "",
		},
		{
			input:     []uint{1, 2, 3},
			delimiter: "",
			expect:    "123",
		},
		{
			input:     []uint{1, 2, 3, 1, 2, 3},
			delimiter: ",",
			expect:    "1,2,3,1,2,3",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
	}
}

func TestJoinFloat32(t *testing.T) {
	testCases := []struct {
		input     []float32
		delimiter string
		expect    string
	}{
		{
			input:  []float32{},
			expect: "",
		},
		{
			input:     []float32{1.1, 2.1, 3.1},
			delimiter: "",
			expect:    "1.12.13.1",
		},
		{
			input:     []float32{1.1, 2.1, 3.1, 1.1, 2.1, 3.1},
			delimiter: ",",
			expect:    "1.1,2.1,3.1,1.1,2.1,3.1",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
	}
}

func TestJoinFloat64(t *testing.T) {
	testCases := []struct {
		input     []float64
		delimiter string
		expect    string
	}{
		{
			input:  []float64{},
			expect: "",
		},
		{
			input:     []float64{1.1, 2.1, 3.1},
			delimiter: "",
			expect:    "1.12.13.1",
		},
		{
			input:     []float64{1.1, 2.1, 3.1, 1.1, 2.1, 3.1},
			delimiter: ",",
			expect:    "1.1,2.1,3.1,1.1,2.1,3.1",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
	}
}

func TestJoinString(t *testing.T) {
	testCases := []struct {
		input     []string
		delimiter string
		expect    string
	}{
		{
			input:  []string{},
			expect: "",
		},
		{
			input:     []string{"a", "b", "c"},
			delimiter: "",
			expect:    "abc",
		},
		{
			input:     []string{"a", "b", "c", "d"},
			delimiter: ",",
			expect:    "a,b,c,d",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expect, Join(testCase.input, testCase.delimiter))
	}
}

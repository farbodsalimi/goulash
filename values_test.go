package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValuessMapString(t *testing.T) {
	testCases := []struct {
		input  map[string]string
		expect []string
	}{
		{
			input:  map[string]string{},
			expect: []string{},
		},
		{
			input:  map[string]string{"key": "value"},
			expect: []string{"value"},
		},
		{
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
			expect: []string{"value1", "value2", "value3", "value4"},
		},
		{
			input:  map[string]string{"key1": "value", "key2": "value", "key3": "value"},
			expect: []string{"value", "value", "value"},
		},
	}

	for _, testCase := range testCases {
		assert.ElementsMatch(t, testCase.expect, Values(testCase.input))
	}
}

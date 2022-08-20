package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysMapString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  map[string]string
		expect []string
	}{
		{
			name:   "Empty map",
			input:  map[string]string{},
			expect: []string{},
		},
		{
			name:   "Map with 1 element",
			input:  map[string]string{"key": "value"},
			expect: []string{"key"},
		},
		{
			name: "Map with distinct keys/values",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
			expect: []string{"key1", "key2", "key3", "key4"},
		},
		{
			name:   "Map with identical values",
			input:  map[string]string{"key1": "value", "key2": "value", "key3": "value"},
			expect: []string{"key1", "key2", "key3"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, testCase.expect, Keys(testCase.input))
		})
	}
}

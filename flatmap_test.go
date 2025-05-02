package goulash

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		expect []string
		fn     func(string) []string
	}{
		{
			name:   "Empty slices",
			input:  []string{"it's sunny in", "", "California"},
			expect: []string{"it's", "sunny", "in", "", "California"},
			fn: func(s string) []string {
				return strings.Split(s, " ")
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expect, FlatMap(testCase.input, testCase.fn))
		})
	}
}

package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		condition   bool
		exprIfTrue  string
		exprIfFalse string
		expect      string
	}{
		{
			name:        "Wrong 2+2 answer",
			condition:   (2 + 2) == 5,
			exprIfTrue:  "yup",
			exprIfFalse: "nope",
			expect:      "nope",
		},
		{
			name:        "Correct 2+2 answer",
			condition:   (2 + 2) == 4,
			exprIfTrue:  "yup",
			exprIfFalse: "nope",
			expect:      "yup",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				testCase.expect,
				Ternary(testCase.condition, testCase.exprIfTrue, testCase.exprIfFalse),
			)
		})
	}
}

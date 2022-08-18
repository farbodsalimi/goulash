package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	testCases := []struct {
		condition   bool
		exprIfTrue  string
		exprIfFalse string
		expect      string
	}{
		{
			condition:   (2 + 2) == 5,
			exprIfTrue:  "yup",
			exprIfFalse: "nope",
			expect:      "nope",
		},
		{
			condition:   (2 + 2) == 4,
			exprIfTrue:  "yup",
			exprIfFalse: "nope",
			expect:      "yup",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(
			t,
			testCase.expect,
			Ternary(testCase.condition, testCase.exprIfTrue, testCase.exprIfFalse),
		)
	}
}
package goulash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	var nilInt []int

	testCases := []struct {
		name      string
		input     any
		predicate any
		wantTrue  any
		wantFalse any
	}{
		{
			name:      "PartitionNumbers",
			input:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(n int) bool { return n%2 == 0 },
			wantTrue:  []int{2, 4, 6},
			wantFalse: []int{1, 3, 5},
		},
		{
			name:      "PartitionStrings",
			input:     []string{"hello", "world", "go", "programming"},
			predicate: func(s string) bool { return len(s) > 4 },
			wantTrue:  []string{"hello", "world", "programming"},
			wantFalse: []string{"go"},
		},
		{
			name:      "PartitionEmpty",
			input:     []int{},
			predicate: func(n int) bool { return n%2 == 0 },
			wantTrue:  nilInt,
			wantFalse: nilInt,
		},
		{
			name:      "PartitionAllTrue",
			input:     []int{2, 4, 6, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			wantTrue:  []int{2, 4, 6, 8},
			wantFalse: nilInt,
		},
		{
			name:      "PartitionAllFalse",
			input:     []int{1, 3, 5, 7},
			predicate: func(n int) bool { return n%2 == 0 },
			wantTrue:  nilInt,
			wantFalse: []int{1, 3, 5, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch input := tc.input.(type) {
			case []int:
				pred := tc.predicate.(func(int) bool)
				gotTrue, gotFalse := Partition(input, pred)
				assert.Equal(t, tc.wantTrue, gotTrue)
				assert.Equal(t, tc.wantFalse, gotFalse)
			case []string:
				pred := tc.predicate.(func(string) bool)
				gotTrue, gotFalse := Partition(input, pred)
				assert.Equal(t, tc.wantTrue, gotTrue)
				assert.Equal(t, tc.wantFalse, gotFalse)
			}
		})
	}
}

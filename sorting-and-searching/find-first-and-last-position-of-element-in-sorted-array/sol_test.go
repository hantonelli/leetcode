package findfirstandlastpositionofelementinsortedarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   10,
			expected: []int{5, 5},
		},
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{8},
			target:   8,
			expected: []int{0, 0},
		},
		{
			nums:     []int{1},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{2, 2},
			target:   2,
			expected: []int{0, 1},
		},
		{
			nums:     []int{1, 3},
			target:   1,
			expected: []int{0, 0},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, searchRange(tt.nums, tt.target))
	}
}

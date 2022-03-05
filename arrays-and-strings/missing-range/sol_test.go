package missingrange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		lower    int
		upper    int
		expected []string
	}{
		{
			nums:     []int{0, 1, 3, 50, 75},
			lower:    0,
			upper:    99,
			expected: []string{"2", "4->49", "51->74", "76->99"},
		},
		{
			nums:     []int{1, 3},
			lower:    0,
			upper:    4,
			expected: []string{"0", "2", "4"},
		},
		{
			nums:     []int{2, 7},
			lower:    0,
			upper:    9,
			expected: []string{"0->1", "3->6", "8->9"},
		},
		{
			nums:     []int{3, 7},
			lower:    0,
			upper:    10,
			expected: []string{"0->2", "4->6", "8->10"},
		},
		{
			nums:     []int{},
			lower:    1,
			upper:    1,
			expected: []string{"1"},
		},
		{
			nums:     []int{},
			lower:    0,
			upper:    1,
			expected: []string{"0->1"},
		},
		{
			nums:     []int{-1},
			lower:    -1,
			upper:    -1,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, findMissingRanges(tt.nums, tt.lower, tt.upper))
	}
}

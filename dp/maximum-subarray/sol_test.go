package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			nums:     []int{1, 2},
			expected: 3,
		},
		{
			nums:     []int{-8, -1, -2},
			expected: -1,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maxSubArrayFast(tt.nums))
	}
}

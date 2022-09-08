package sum3closest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			nums:     []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{0, 1, 2},
			target:   3,
			expected: 3,
		},
		{
			nums:     []int{1, 1, 1, 0},
			target:   100,
			expected: 3,
		},
		{
			nums:     []int{0, 2, 1, -3},
			target:   1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, threeSumClosest(tt.nums, tt.target))
	}
}

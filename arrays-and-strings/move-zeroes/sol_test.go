package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		moveZeroes(tt.nums)
		require.EqualValues(t, tt.expected, tt.nums)
	}
}

package canjump

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			nums:     []int{0},
			expected: true,
		},
		{
			nums:     []int{2, 0},
			expected: true,
		},
		{
			nums:     []int{2, 0, 0},
			expected: true,
		},
		{
			nums:     []int{2, 0, 0, 1, 1},
			expected: false,
		},
		{
			nums:     []int{2, 0, 1, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, canJump(tt.nums))
	}
}

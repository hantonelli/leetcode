package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			nums:     []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			nums:     []int{2, 1},
			k:        4,
			expected: 1,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, findKthLargest(tt.nums, tt.k))
	}

}

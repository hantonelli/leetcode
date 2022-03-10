package top_k_frequent_elements

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, topKFrequent(tt.nums, tt.k))
	}
}

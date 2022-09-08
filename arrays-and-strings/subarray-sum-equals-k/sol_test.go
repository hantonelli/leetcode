package subarray_sum_equals_k

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
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		// {
		// 	nums:     []int{1, 2, 3},
		// 	k:        3,
		// 	expected: 2,
		// },
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, subarraySum(tt.nums, tt.k))
	}
}

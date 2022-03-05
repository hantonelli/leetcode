package nextpermutation

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
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
	}

	for _, tt := range tests {
		nextPermutation(tt.nums)
		require.EqualValues(t, tt.expected, tt.nums)
	}
}

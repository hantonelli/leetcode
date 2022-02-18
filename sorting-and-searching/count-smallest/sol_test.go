package countsmallest

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
			nums:     []int{5, 2, 6, 1},
			expected: []int{2, 1, 1, 0},
		},
		{
			nums:     []int{-1},
			expected: []int{0},
		},
		{
			nums:     []int{-1, -1},
			expected: []int{0, 0},
		},
	}
	for _, tt := range tests {
		require.Equal(t, tt.expected, countSmaller(tt.nums))
	}
}

package houseofrobber

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
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, rob2(tt.nums))
	}
}

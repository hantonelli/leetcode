package mincostclimbingstairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		cost     []int
		expected int
	}{
		// {
		// 	cost:     []int{10, 15, 20},
		// 	expected: 15,
		// },
		{
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, minCostClimbingStairs2(tt.cost))
	}
}

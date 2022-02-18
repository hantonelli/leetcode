package coinchange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20,
		},
		{
			coins:    []int{186, 419, 83, 408},
			amount:   6249,
			expected: 20,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, coinChange2(tt.coins, tt.amount))
	}
}

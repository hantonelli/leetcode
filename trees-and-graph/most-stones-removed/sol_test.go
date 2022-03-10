package most_stones_removed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		stones   [][]int
		expected int
	}{
		{
			stones: [][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 2},
				{2, 1},
				{2, 2},
			},
			expected: 5,
		},
		{
			stones: [][]int{
				{2, 2},
			},
			expected: 0,
		},
		{
			stones: [][]int{
				{0, 0},
				{2, 2},
			},
			expected: 0,
		},
		{
			stones: [][]int{
				{1, 2},
				{2, 2},
			},
			expected: 1,
		},
		{
			stones:   [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 3}, {2, 3}, {0, 2}},
			expected: 5,
		},
		{
			stones: [][]int{
				{5, 9}, {9, 0}, {0, 0}, {7, 0}, {4, 3}, {8, 5}, {5, 8}, {1, 1}, {0, 6}, {7, 5}, {1, 6}, {1, 9}, {9, 4}, {2, 8}, {1, 3}, {4, 2}, {2, 5}, {4, 1}, {0, 2}, {6, 5},
			},
			expected: 19,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, removeStones(tt.stones))
	}
}

package validatebinarysearchtree

import (
	"fmt"
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		tree     []int
		expected bool
	}{
		{
			tree:     []int{2, 1, 3},
			expected: true,
		},
		{
			tree:     []int{5, 1, 4, ds.NULL, ds.NULL, 3, 6},
			expected: false,
		},
		{
			tree:     []int{5, 4, 8, ds.NULL, 6},
			expected: false,
		},
		{
			tree:     []int{5, 4, 6, ds.NULL, ds.NULL, 3, 7},
			expected: false,
		},
		{
			tree:     []int{5, 3, 8, ds.NULL, ds.NULL, 7, 9, 4},
			expected: false,
		},
		{
			tree:     []int{5, 4, 8, ds.NULL, ds.NULL, 7, 9, 6},
			expected: true,
		},
		{
			tree:     []int{2, 2, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, isValidBST(ds.ListToTree(tt.tree)), fmt.Sprint(tt.tree))
	}
}

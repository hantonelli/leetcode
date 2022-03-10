package binary_tree_zigzag_level_order_traversal

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		root     *ds.TreeNode
		expected [][]int
	}{
		{
			root:     ds.ListToTree([]int{3, 9, 20, ds.NULL, ds.NULL, 15, 7}),
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root:     ds.ListToTree([]int{1}),
			expected: [][]int{{1}},
		},
		{
			root:     ds.ListToTree([]int{1, 2, 3, 4, ds.NULL, ds.NULL, 5}),
			expected: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, zigzagLevelOrder(tt.root))
	}
}

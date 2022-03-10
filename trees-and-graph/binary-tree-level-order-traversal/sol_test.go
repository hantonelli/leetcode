package binary_tree_level_order_traversal

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
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			root:     ds.ListToTree([]int{1}),
			expected: [][]int{{1}},
		},
		{
			root:     ds.ListToTree([]int{}),
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, levelOrder(tt.root))
	}
}

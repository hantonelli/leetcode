package flatten_binary_tree_to_linked_list

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		root   *ds.TreeNode
		expect []int
	}{
		{
			root:   ds.ListToTree([]int{1, 2, 5, 3, 4, ds.NULL, 6}),
			expect: []int{1, ds.NULL, 2, ds.NULL, 3, ds.NULL, 4, ds.NULL, 5, ds.NULL, 6},
		},
		{
			root:   ds.ListToTree([]int{}),
			expect: []int{},
		},
		{
			root:   ds.ListToTree([]int{0}),
			expect: []int{0},
		},
	}

	for _, tt := range tests {
		flatten(tt.root)
		require.Equal(t, tt.expect, ds.TreetoList(tt.root))
	}
}

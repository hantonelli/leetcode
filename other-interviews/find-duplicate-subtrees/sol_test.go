package findduplicatesubtrees

import (
	"testing"

	ds "github.com/hantonelli/leetcode/data-structures"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	root := ds.ListToTree([]int{1, 2, 3, 4, ds.NULL, 2, 4, ds.NULL, ds.NULL, 4})
	expected := []*ds.TreeNode{root.Right.Left.Left, root.Right.Left}
	require.Equal(t, expected, findDuplicateSubtrees(root))
}

func Test2(t *testing.T) {
	root := ds.ListToTree([]int{2, 1, 1})
	expected := []*ds.TreeNode{root.Right}
	require.Equal(t, expected, findDuplicateSubtrees(root))
}

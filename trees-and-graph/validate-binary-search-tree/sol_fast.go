package validatebinarysearchtree

import "github.com/hantonelli/leetcode/ds"

func isValidBSTFast(root *ds.TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node, low, high *ds.TreeNode) bool {
	if node == nil {
		return true
	}
	if (low != nil && low.Val >= node.Val) || (high != nil && high.Val <= node.Val) {
		return false
	}
	return helper(node.Left, low, node) && helper(node.Right, node, high)

}

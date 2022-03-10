package symmetrictree

import "github.com/hantonelli/leetcode/ds"

func isSymmetric(root *ds.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricInner(root.Left, root.Right)
}

func isSymmetricInner(left, right *ds.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil || left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricInner(left.Left, right.Right) && isSymmetricInner(left.Right, right.Left)
}

package validatebinarysearchtree

import "github.com/hantonelli/leetcode/ds"

func isValidBST(root *ds.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}
		if leftValid := isValidBST(root.Left); !leftValid {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if rightValid := isValidBST(root.Right); !rightValid {
			return false
		}
	}
	return true
}

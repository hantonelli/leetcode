package validatebinarysearchtree

import (
	"math"

	"github.com/hantonelli/leetcode/ds"
)

func isValidBST(r *ds.TreeNode) bool {
	if r == nil {
		return false
	}
	return isValidBSTInner(r.Left, math.MinInt64, r.Val) &&
		isValidBSTInner(r.Right, r.Val, math.MaxInt64)
}

func isValidBSTInner(r *ds.TreeNode, minV, maxV int) bool {
	if r == nil {
		return true
	}
	if r.Val <= minV || maxV <= r.Val {
		return false
	}
	if r.Left != nil {
		if leftValid := isValidBSTInner(r.Left, minV, r.Val); !leftValid {
			return false
		}
	}
	if r.Right != nil {
		if rightValid := isValidBSTInner(r.Right, r.Val, maxV); !rightValid {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

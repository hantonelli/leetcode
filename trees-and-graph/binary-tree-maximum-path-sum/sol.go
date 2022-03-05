package main

import (
	"math"

	"github.com/hantonelli/leetcode/ds"
)

func maxPathSum(n *ds.TreeNode) int {
	maxHalf, maxFull := maxPathSumFull(n)
	return ds.Max(maxHalf, maxFull)
}

func maxPathSumFull(n *ds.TreeNode) (maxHalf int, maxFull int) {
	if n == nil {
		return 0, 0
	}

	rightHalf, rightFull := math.MinInt64, math.MinInt64
	if n.Right != nil {
		rightHalf, rightFull = maxPathSumFull(n.Right)
	}
	leftHalf, leftFull := math.MinInt64, math.MinInt64
	if n.Left != nil {
		leftHalf, leftFull = maxPathSumFull(n.Left)
	}

	maxHalf = n.Val
	if 0 < leftHalf || 0 < rightHalf {
		if leftHalf < rightHalf {
			maxHalf += rightHalf
		} else {
			maxHalf += leftHalf
		}
	}

	maxFull = n.Val
	if n.Right != nil && 0 < rightHalf {
		maxFull += rightHalf
	}
	if n.Left != nil && 0 < leftHalf {
		maxFull += leftHalf
	}

	return maxHalf, ds.Max(maxFull, ds.Max(rightFull, leftFull))
}

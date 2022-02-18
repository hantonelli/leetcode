package main

import (
	ds "github.com/hantonelli/leetcode/data-structures"
)

func diameterOfBinaryTree(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}
	lDist, lMax := getDistance(root.Left)
	rDist, rMax := getDistance(root.Right)
	dist := max(lDist, rDist)
	maxD := max(lDist+rDist, max(lMax, rMax))
	return max(dist, maxD)
}

func getDistance(n *ds.TreeNode) (int, int) {
	if n == nil {
		return 0, 0
	}
	lDist, lMax := getDistance(n.Left)
	rDist, rMax := getDistance(n.Right)
	dist := max(lDist+1, rDist+1)
	max := max(lDist+rDist, max(lMax, rMax))
	return dist, max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

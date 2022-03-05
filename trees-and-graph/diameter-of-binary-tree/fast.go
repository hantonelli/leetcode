package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func diameterOfBinaryTreeF(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}
	lDist, lMax := getDistance(root.Left)
	rDist, rMax := getDistance(root.Right)
	return max(max(lDist, rDist), max(lDist+rDist, max(lMax, rMax)))
}

func getDistanceF(n *ds.TreeNode) (int, int) {
	if n == nil {
		return 0, 0
	}
	lDist, lMax := getDistance(n.Left)
	rDist, rMax := getDistance(n.Right)
	return max(lDist+1, rDist+1), max(lDist+rDist, max(lMax, rMax))
}

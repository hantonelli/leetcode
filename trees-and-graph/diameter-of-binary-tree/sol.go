package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDist, lMax := getDistance(root.Left)
	rDist, rMax := getDistance(root.Right)
	dist := math.Max(lDist, rDist)
	max := math.Max(lMax, rMax)
	return int(math.Max(dist, max))
}

func getDistance(n *TreeNode) (float64, float64) {
	if n == nil {
		return 1, 0
	}
	lDist, lMax := getDistance(n.Left)
	rDist, rMax := getDistance(n.Right)
	return math.Max(lDist, rDist), math.Max(lMax, rMax)
}

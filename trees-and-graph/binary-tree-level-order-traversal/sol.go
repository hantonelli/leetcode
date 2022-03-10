package binary_tree_level_order_traversal

import "github.com/hantonelli/leetcode/ds"

func levelOrder(root *ds.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	currLevel := []*ds.TreeNode{root}

	for level := 0; 0 < len(currLevel); level++ {
		nextLevel := []*ds.TreeNode{}
		levelVals := []int{}
		for _, n := range currLevel {
			levelVals = append(levelVals, n.Val)
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
		res = append(res, levelVals)
		currLevel = nextLevel
	}
	return res
}

package binary_tree_zigzag_level_order_traversal

import "github.com/hantonelli/leetcode/ds"

func zigzagLevelOrder(root *ds.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	curr := []*ds.TreeNode{root}
	leftToRight := true

	for level := 0; 0 < len(curr); level++ {
		nextLevel := []*ds.TreeNode{}
		levelVals := []int{}
		if leftToRight {
			for i := len(curr) - 1; 0 <= i; i-- {
				levelVals = append(levelVals, curr[i].Val)
				if curr[i].Left != nil {
					nextLevel = append(nextLevel, curr[i].Left)
				}
				if curr[i].Right != nil {
					nextLevel = append(nextLevel, curr[i].Right)
				}
			}
		} else {
			for i := len(curr) - 1; 0 <= i; i-- {
				levelVals = append(levelVals, curr[i].Val)
				if curr[i].Right != nil {
					nextLevel = append(nextLevel, curr[i].Right)
				}
				if curr[i].Left != nil {
					nextLevel = append(nextLevel, curr[i].Left)
				}
			}
		}
		res = append(res, levelVals)
		curr = nextLevel
		leftToRight = !leftToRight
	}
	return res
}

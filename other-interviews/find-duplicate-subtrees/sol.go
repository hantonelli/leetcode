package findduplicatesubtrees

import (
	"fmt"

	"github.com/hantonelli/leetcode/ds"
)

func findDuplicateSubtrees(root *ds.TreeNode) []*ds.TreeNode {
	duplicated := []*ds.TreeNode{}
	viewed := map[string]int{}
	isDuplicate(root, &duplicated, viewed)
	return duplicated
}

func isDuplicate(root *ds.TreeNode, duplicated *[]*ds.TreeNode, viewed map[string]int) []int {
	if root == nil {
		return nil
	}

	nodes := []int{root.Val}
	if root.Left != nil || root.Right != nil {
		if root.Left == nil {
			nodes = append(nodes, ds.NULL)
		} else {
			leftNodes := isDuplicate(root.Left, duplicated, viewed)
			if leftNodes == nil {
				nodes = append(nodes, ds.NULL)
			} else {
				nodes = append(nodes, leftNodes...)
			}
		}
		if root.Right == nil {
			nodes = append(nodes, ds.NULL)
		} else {
			rightNodes := isDuplicate(root.Right, duplicated, viewed)
			if rightNodes == nil {
				nodes = append(nodes, ds.NULL)
			} else {
				nodes = append(nodes, rightNodes...)
			}
		}
	}
	if count, ok := viewed[fmt.Sprint(nodes)]; ok && count == 1 {
		*duplicated = append(*duplicated, root)
	}
	viewed[fmt.Sprint(nodes)]++
	return nodes
}

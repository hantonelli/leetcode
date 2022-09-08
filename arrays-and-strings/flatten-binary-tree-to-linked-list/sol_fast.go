package flatten_binary_tree_to_linked_list

import "github.com/hantonelli/leetcode/ds"

func flattenFast(root *ds.TreeNode) {
	if root == nil {
		return
	}
	node := root
	for node != nil {
		if node.Left != nil {
			rightMost := node.Left
			for rightMost.Right != nil {
				rightMost = rightMost.Right
			}
			//Rewire
			rightMost.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}
		node = node.Right
	}
}

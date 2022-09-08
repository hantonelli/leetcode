package flatten_binary_tree_to_linked_list

import "github.com/hantonelli/leetcode/ds"

func flatten(root *ds.TreeNode) {
	if root == nil {
		return
	}
	flattenInner(root)
	curr := root
	for curr != nil {
		curr.Right = curr.Left
		curr.Left = nil
		curr = curr.Right
	}
}

func flattenInner(node *ds.TreeNode) *ds.TreeNode {
	var lastNode *ds.TreeNode
	if node.Left != nil {
		lastNode = flattenInner(node.Left)
	} else {
		if node.Right != nil {
			node.Left = node.Right
			node.Right = nil
			lastNode = flattenInner(node.Left)
		}
	}
	if lastNode != nil && node.Right != nil {
		lastNode.Left = node.Right
		node.Right = nil
		lastNode = flattenInner(node.Left)
	}
	if lastNode != nil {
		return lastNode
	}
	return node
}

package binarytreepaths

import (
	"fmt"

	"github.com/hantonelli/leetcode/ds"
)

func binaryTreePaths(r *ds.TreeNode) []string {
	if r == nil {
		return nil
	}
	if r.Left == nil && r.Right == nil {
		return []string{fmt.Sprintf("%d", r.Val)}
	}

	res := []string{}
	if r.Left != nil {
		childs := binaryTreePaths(r.Left)
		for _, v := range childs {
			res = append(res, fmt.Sprintf("%d->%s", r.Val, v))
		}
	}
	if r.Right != nil {
		childs := binaryTreePaths(r.Right)
		for _, v := range childs {
			res = append(res, fmt.Sprintf("%d->%s", r.Val, v))
		}
	}
	return res
}

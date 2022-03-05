package findduplicatesubtrees

import (
	"fmt"

	ds "github.com/hantonelli/leetcode/data-structures"
)

func findDuplicateSubtreesFast(root *ds.TreeNode) []*ds.TreeNode {
	var res []*ds.TreeNode
	id := 1
	serialToId := map[string]int{}
	idToCount := map[int]int{}
	postOrder(root, serialToId, idToCount, &id, &res)
	return res
}

func postOrder(root *ds.TreeNode, serialToId map[string]int, idToCount map[int]int, id *int, res *[]*ds.TreeNode) int {
	if root == nil {
		return 0
	}
	left := postOrder(root.Left, serialToId, idToCount, id, res)
	right := postOrder(root.Right, serialToId, idToCount, id, res)
	serial := fmt.Sprintf("%d,%d,%d", left, root.Val, right)
	serialId, exist := serialToId[serial]
	if !exist {
		serialId = *id
		serialToId[serial] = *id
		*id++
	}
	idToCount[serialId]++
	if idToCount[serialId] == 2 {
		*res = append(*res, root)
	}
	return serialId
}

package reorder_list

import "github.com/hantonelli/leetcode/ds"

func reorderList(head *ds.ListNode) {
	nodes := []*ds.ListNode{}
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	prev := head
	fromStart := false
	start := 1
	end := len(nodes) - 1
	for i := 1; i < len(nodes); i++ {
		if fromStart {
			prev.Next = nodes[start]
			prev = nodes[start]
			start++
		} else {
			prev.Next = nodes[end]
			prev = nodes[end]
			end--
		}
		fromStart = !fromStart
	}
	prev.Next = nil
}

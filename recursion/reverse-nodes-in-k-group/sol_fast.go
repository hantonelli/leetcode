package reverse_nodes_in_k_group

import "github.com/hantonelli/leetcode/ds"

func reverseKGroupFast(head *ds.ListNode, k int) *ds.ListNode {
	needSetHead := true
	current := head
	var lastPrev, sublistStart *ds.ListNode
	for {
		sublistStart = current
		fullPath := true
		for i := 0; i < k; i++ {
			if i == k-1 {
				if lastPrev != nil {
					lastPrev.Next = current
				}

				if needSetHead {
					head = current
					needSetHead = false
				}
			}
			if current == nil {
				if i-1 != k-1 {
					fullPath = false
				}
				break
			}
			current = current.Next
		}
		if current == nil && !fullPath {
			if lastPrev == nil {
				return head
			}
			lastPrev.Next = sublistStart
			break
		}

		prev := sublistStart
		sublistStart = sublistStart.Next
		prev.Next = nil
		lastPrev = prev
		for {
			if sublistStart == current {
				break
			}
			tmp := sublistStart.Next
			sublistStart.Next = prev
			prev = sublistStart
			sublistStart = tmp
		}
	}
	return head
}

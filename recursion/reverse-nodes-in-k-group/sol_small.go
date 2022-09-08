package reverse_nodes_in_k_group

import "github.com/hantonelli/leetcode/ds"

func reverseKGroupSmall(head *ds.ListNode, k int) *ds.ListNode {
	if k <= 1 {
		return head
	}

	var newHead *ds.ListNode
	var newTail *ds.ListNode
	ptr := head

	for ptr != nil {
		c := k
		start := ptr
		for c > 0 && ptr != nil {
			ptr = ptr.Next
			c--
		}
		if c == 0 {
			h, t := reverseK(start, k)
			if newHead == nil {
				newHead = h
				newTail = t
			} else {
				newTail.Next = h
				newTail = t
			}
		} else {
			newTail.Next = start
		}
	}
	return newHead
}

func reverseK(head *ds.ListNode, k int) (newHead *ds.ListNode, tail *ds.ListNode) {
	ptr := head
	idx := k
	for idx > 0 {
		n := ptr.Next
		ptr.Next = newHead
		newHead = ptr
		ptr = n
		if tail == nil {
			tail = newHead
		}
		idx--
	}
	return
}

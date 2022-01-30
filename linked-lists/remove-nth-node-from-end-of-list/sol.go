package main

import (
	ll "github.com/hantonelli/leetcode/data-structures"
)

func removeNthFromEnd(head *ll.ListNode, n int) *ll.ListNode {
	if n < 0 {
		return head
	}
	var nNode, nNodePrev *ll.ListNode
	curr := head
	count := 1
	for curr != nil {
		if nNode != nil {
			nNodePrev = nNode
			nNode = nNode.Next
		}
		if count == n {
			nNode = head
		}
		count++
		curr = curr.Next
	}
	if nNodePrev != nil {
		if nNode.Next != nil {
			nNodePrev.Next = nNode.Next
		} else {
			nNodePrev.Next = nil
		}
		return head
	}
	if nNode != nil {
		return nNode.Next
	}
	return nil
}

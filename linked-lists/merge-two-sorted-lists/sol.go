package main

import (
	ll "github.com/hantonelli/leetcode/data-structures"
)

func mergeTwoLists(list1, list2 *ll.ListNode) *ll.ListNode {
	root := &ll.ListNode{}
	prev := root
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
		prev.Next = nil
	}
	for list1 != nil {
		prev.Next = list1
		list1 = list1.Next
		prev = prev.Next
		prev.Next = nil
	}
	for list2 != nil {
		prev.Next = list2
		list2 = list2.Next
		prev = prev.Next
		prev.Next = nil
	}
	return root.Next
}
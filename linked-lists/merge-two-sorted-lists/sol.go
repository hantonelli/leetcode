package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func mergeTwoLists(list1, list2 *ds.ListNode) *ds.ListNode {
	root := &ds.ListNode{}
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

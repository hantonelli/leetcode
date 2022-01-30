package main

import (
	ll "github.com/hantonelli/leetcode/data-structures"
)

func mergeKLists(list []*ll.ListNode) *ll.ListNode {
	var root, prev, curr *ll.ListNode
	var listsLowest int
	for {
		listsLowest = -1
		for i := range list {
			if list[i] != nil && (listsLowest == -1 || list[i].Val < list[listsLowest].Val) {
				listsLowest = i
			}
		}
		if listsLowest == -1 {
			break
		}
		curr = &ll.ListNode{
			Val: list[listsLowest].Val,
		}
		if prev == nil {
			root = curr
		} else {
			prev.Next = curr
		}
		list[listsLowest] = list[listsLowest].Next
		prev = curr
	}

	return root
}

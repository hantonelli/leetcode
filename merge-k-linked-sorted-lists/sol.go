package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func mergeKLists(list []*ds.ListNode) *ds.ListNode {
	var root, prev, curr *ds.ListNode
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
		curr = &ds.ListNode{
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

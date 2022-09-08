package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func mergeKListsF(data []*ds.ListNode) *ds.ListNode {
	if len(data) == 0 {
		return nil
	}
	if len(data) == 1 {
		return data[0]
	}
	middle := len(data) / 2
	left := mergeKListsF(data[:middle])
	right := mergeKListsF(data[middle:])

	return mergeF(left, right)
}

func mergeF(l, r *ds.ListNode) *ds.ListNode {
	root := &ds.ListNode{}
	prev := root
	for l != nil && r != nil {
		if l.Val < r.Val {
			prev.Next = l
			l = l.Next
		} else {
			prev.Next = r
			r = r.Next
		}
		prev = prev.Next
		prev.Next = nil
	}
	for l != nil {
		prev.Next = l
		l = l.Next
		prev = prev.Next
		prev.Next = nil
	}
	for r != nil {
		prev.Next = r
		r = r.Next
		prev = prev.Next
		prev.Next = nil
	}
	return root.Next
}

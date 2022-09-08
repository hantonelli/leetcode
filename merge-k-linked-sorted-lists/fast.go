package main

import "github.com/hantonelli/leetcode/ds"

func mergeKListsFast(lists []*ds.ListNode) *ds.ListNode {

	n := len(lists)
	if n == 0 {
		return nil
	}
	return mergeSort(lists, 0, n-1)
}

func mergeSort(lists []*ds.ListNode, a int, z int) *ds.ListNode {
	if a == z {
		return lists[a]
	}
	m := a + ((z - a) / 2)
	x := mergeSort(lists, a, m)
	y := mergeSort(lists, m+1, z)
	return merge(x, y)
}

func merge(x *ds.ListNode, y *ds.ListNode) *ds.ListNode {
	node := ds.ListNode{}
	head := &node
	ptr := head
	for x != nil || y != nil {
		if x != nil && y != nil {
			if x.Val < y.Val {
				ptr.Next = x
				x = x.Next
				ptr = ptr.Next
			} else {
				ptr.Next = y
				y = y.Next
				ptr = ptr.Next
			}
		} else if x != nil {
			ptr.Next = x
			x = x.Next
			ptr = ptr.Next
		} else {
			ptr.Next = y
			y = y.Next
			ptr = ptr.Next
		}
	}
	return head.Next
}

package main

import ll "github.com/hantonelli/leetcode/data-structures"

func mergeKListsFast(lists []*ll.ListNode) *ll.ListNode {

	n := len(lists)
	if n == 0 {
		return nil
	}
	return mergeSort(lists, 0, n-1)
}

func mergeSort(lists []*ll.ListNode, a int, z int) *ll.ListNode {
	if a == z {
		return lists[a]
	}
	m := a + ((z - a) / 2)
	x := mergeSort(lists, a, m)
	y := mergeSort(lists, m+1, z)
	return merge(x, y)
}

func merge(x *ll.ListNode, y *ll.ListNode) *ll.ListNode {
	node := ll.ListNode{}
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

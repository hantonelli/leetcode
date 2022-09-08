package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	head := ds.ListNode{}
	curr := &head
	sum := 0
	rest := 0
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += rest
		curr.Next = &ds.ListNode{Val: sum % 10}
		curr = curr.Next
		rest = int(sum / 10)
	}
	if rest != 0 {
		curr.Next = &ds.ListNode{Val: rest}
	}
	return head.Next
}

func addTwoNumbers1(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	head := ds.ListNode{}
	curr := &head
	rest := 0
	for l1 != nil && l2 != nil {
		num := l1.Val + l2.Val + rest
		curr.Next = &ds.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		num := l1.Val + rest
		curr.Next = &ds.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l1 = l1.Next
	}
	for l2 != nil {
		num := l2.Val + rest
		curr.Next = &ds.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l2 = l2.Next
	}
	if rest != 0 {
		curr.Next = &ds.ListNode{
			Val: rest,
		}
	}
	return head.Next
}

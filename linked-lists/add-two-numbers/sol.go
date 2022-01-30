package main

import (
	ll "github.com/hantonelli/leetcode/data-structures"
)

func addTwoNumbers(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	head := ll.ListNode{}
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
		curr.Next = &ll.ListNode{Val: sum % 10}
		curr = curr.Next
		rest = int(sum / 10)
	}
	if rest != 0 {
		curr.Next = &ll.ListNode{Val: rest}
	}
	return head.Next
}

func addTwoNumbers1(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	head := ll.ListNode{}
	curr := &head
	rest := 0
	for l1 != nil && l2 != nil {
		num := l1.Val + l2.Val + rest
		curr.Next = &ll.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		num := l1.Val + rest
		curr.Next = &ll.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l1 = l1.Next
	}
	for l2 != nil {
		num := l2.Val + rest
		curr.Next = &ll.ListNode{
			Val: num % 10,
		}
		curr = curr.Next
		rest = int(num / 10)
		l2 = l2.Next
	}
	if rest != 0 {
		curr.Next = &ll.ListNode{
			Val: rest,
		}
	}
	return head.Next
}

package dataStructures

import "fmt"

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l.Next != nil {
		return fmt.Sprintf("%d, %s", l.Val, l.Next.String())
	}
	return fmt.Sprintf("%d", l.Val)
}

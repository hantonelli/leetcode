package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	ll "github.com/hantonelli/leetcode/data-structures"
)

func Test(t *testing.T) {
	tests := []struct {
		l1       *ll.ListNode
		l2       *ll.ListNode
		expected *ll.ListNode
	}{
		{
			l1:       ll.ArrToLinkedList([]int{2, 4, 3}),
			l2:       ll.ArrToLinkedList([]int{5, 6, 4}),
			expected: ll.ArrToLinkedList([]int{7, 0, 8}),
		},
		{
			l1:       ll.ArrToLinkedList([]int{0}),
			l2:       ll.ArrToLinkedList([]int{0}),
			expected: ll.ArrToLinkedList([]int{0}),
		},
		{
			l1:       ll.ArrToLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       ll.ArrToLinkedList([]int{9, 9, 9, 9}),
			expected: ll.ArrToLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		res := addTwoNumbers(tt.l1, tt.l2)
		require.Equal(t, tt.expected, res)
	}
}

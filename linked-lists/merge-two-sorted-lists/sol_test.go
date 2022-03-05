package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	ll "github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		list1    *ll.ListNode
		list2    *ll.ListNode
		expected *ll.ListNode
	}{
		{
			list1:    ll.ArrToLinkedList([]int{1, 2, 4}),
			list2:    ll.ArrToLinkedList([]int{1, 3, 4}),
			expected: ll.ArrToLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			list1:    ll.ArrToLinkedList([]int{}),
			list2:    ll.ArrToLinkedList([]int{}),
			expected: ll.ArrToLinkedList([]int{}),
		},
		{
			list1:    ll.ArrToLinkedList([]int{}),
			list2:    ll.ArrToLinkedList([]int{0}),
			expected: ll.ArrToLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		res := mergeTwoLists(tt.list1, tt.list2)
		require.Equal(t, tt.expected, res)
	}
}

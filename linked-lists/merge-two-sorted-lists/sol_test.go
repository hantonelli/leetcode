package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		list1    *ds.ListNode
		list2    *ds.ListNode
		expected *ds.ListNode
	}{
		{
			list1:    ds.ArrToLinkedList([]int{1, 2, 4}),
			list2:    ds.ArrToLinkedList([]int{1, 3, 4}),
			expected: ds.ArrToLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			list1:    ds.ArrToLinkedList([]int{}),
			list2:    ds.ArrToLinkedList([]int{}),
			expected: ds.ArrToLinkedList([]int{}),
		},
		{
			list1:    ds.ArrToLinkedList([]int{}),
			list2:    ds.ArrToLinkedList([]int{0}),
			expected: ds.ArrToLinkedList([]int{0}),
		},
	}

	for _, tt := range tests {
		res := mergeTwoLists(tt.list1, tt.list2)
		require.Equal(t, tt.expected, res)
	}
}

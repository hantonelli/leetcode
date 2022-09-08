package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		l1       *ds.ListNode
		l2       *ds.ListNode
		expected *ds.ListNode
	}{
		{
			l1:       ds.ArrToLinkedList([]int{2, 4, 3}),
			l2:       ds.ArrToLinkedList([]int{5, 6, 4}),
			expected: ds.ArrToLinkedList([]int{7, 0, 8}),
		},
		{
			l1:       ds.ArrToLinkedList([]int{0}),
			l2:       ds.ArrToLinkedList([]int{0}),
			expected: ds.ArrToLinkedList([]int{0}),
		},
		{
			l1:       ds.ArrToLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       ds.ArrToLinkedList([]int{9, 9, 9, 9}),
			expected: ds.ArrToLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		res := addTwoNumbers(tt.l1, tt.l2)
		require.Equal(t, tt.expected, res)
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		l        *ds.ListNode
		n        int
		expected *ds.ListNode
	}{
		{
			l:        ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        2,
			expected: ds.ArrToLinkedList([]int{1, 2, 3, 5}),
		},
		{
			l:        ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        1,
			expected: ds.ArrToLinkedList([]int{1, 2, 3, 4}),
		},
		{
			l:        ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        5,
			expected: ds.ArrToLinkedList([]int{2, 3, 4, 5}),
		},
		{
			l:        ds.ArrToLinkedList([]int{1}),
			n:        1,
			expected: ds.ArrToLinkedList([]int{}),
		},
		{
			l:        ds.ArrToLinkedList([]int{1, 2}),
			n:        1,
			expected: ds.ArrToLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		res := removeNthFromEnd(tt.l, tt.n)
		require.Equal(t, tt.expected, res)
	}
}

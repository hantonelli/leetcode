package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	ll "github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		l        *ll.ListNode
		n        int
		expected *ll.ListNode
	}{
		{
			l:        ll.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        2,
			expected: ll.ArrToLinkedList([]int{1, 2, 3, 5}),
		},
		{
			l:        ll.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        1,
			expected: ll.ArrToLinkedList([]int{1, 2, 3, 4}),
		},
		{
			l:        ll.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			n:        5,
			expected: ll.ArrToLinkedList([]int{2, 3, 4, 5}),
		},
		{
			l:        ll.ArrToLinkedList([]int{1}),
			n:        1,
			expected: ll.ArrToLinkedList([]int{}),
		},
		{
			l:        ll.ArrToLinkedList([]int{1, 2}),
			n:        1,
			expected: ll.ArrToLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		res := removeNthFromEnd(tt.l, tt.n)
		require.Equal(t, tt.expected, res)
	}
}

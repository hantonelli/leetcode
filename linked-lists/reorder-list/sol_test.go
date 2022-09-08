package reorder_list

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		head     *ds.ListNode
		expected *ds.ListNode
	}{
		{
			head:     ds.ArrToLinkedList([]int{1, 2, 3, 4}),
			expected: ds.ArrToLinkedList([]int{1, 4, 2, 3}),
		},
		{
			head:     ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			expected: ds.ArrToLinkedList([]int{1, 5, 2, 4, 3}),
		},
	}

	for _, tt := range tests {
		reorderList(tt.head)
		require.Equal(t, tt.expected, tt.head)
	}
}

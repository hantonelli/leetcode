package reverse_linked_list

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
			head:     ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			expected: ds.ArrToLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			head:     ds.ArrToLinkedList([]int{1}),
			expected: ds.ArrToLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, reverseList(tt.head))
	}
}

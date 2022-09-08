package reverse_nodes_in_k_group

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		head     *ds.ListNode
		k        int
		expected *ds.ListNode
	}{
		{
			head:     ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			k:        2,
			expected: ds.ArrToLinkedList([]int{2, 1, 4, 3, 5}),
		},
		{
			head:     ds.ArrToLinkedList([]int{1, 2, 3, 4, 5}),
			k:        3,
			expected: ds.ArrToLinkedList([]int{3, 2, 1, 4, 5}),
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, reverseKGroupFast(tt.head, tt.k))
	}
}

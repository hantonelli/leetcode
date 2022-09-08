package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hantonelli/leetcode/ds"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []*ds.ListNode
		expected *ds.ListNode
	}{
		{
			input:    ds.ArrOfArrToArrLinkedList([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}),
			expected: ds.ArrToLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}

	for _, tt := range tests {
		res := mergeKListsF(tt.input)
		require.Equal(t, tt.expected, res)
	}
}

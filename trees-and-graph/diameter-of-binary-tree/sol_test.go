package main

import (
	"testing"

	ds "github.com/hantonelli/leetcode/data-structures"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    *ds.TreeNode
		expected int
	}{
		{
			input:    ds.ListToTree([]int{1, 2}),
			expected: 1,
		},
		{
			input:    ds.ListToTree([]int{1, 2, 3, 4, 5}),
			expected: 3,
		},
		{
			input:    ds.ListToTree([]int{4, -7, -3, ds.NULL, ds.NULL, -9, -3, 9, -7, -4, ds.NULL, 6, ds.NULL, -6, -6, ds.NULL, ds.NULL, 0, 6, 5, ds.NULL, 9, ds.NULL, ds.NULL, -1, -4, ds.NULL, ds.NULL, ds.NULL, -2}),
			expected: 8,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, diameterOfBinaryTree(tt.input))
	}
}

package symmetrictree

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		root     *ds.TreeNode
		expected bool
	}{
		{
			root:     ds.ListToTree([]int{1, 2, 2, 3, 4, 4, 3}),
			expected: true,
		},
		{
			root:     ds.ListToTree([]int{1, 2, 2, ds.NULL, 3, ds.NULL, 3}),
			expected: false,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, isSymmetric(tt.root))
	}
}

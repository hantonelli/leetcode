package dataStructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeMapper(t *testing.T) {
	tests := []struct {
		list []int
	}{
		{
			list: []int{4, -7, -3, NULL, NULL, -9, -3, 9, -7, -4, NULL, 6, NULL, -6, -6, NULL, NULL, 0, 6, 5, NULL, 9, NULL, NULL, -1, -4, NULL, NULL, NULL, -2},
		},
	}

	for _, tt := range tests {
		tree := ListToTree(tt.list)
		rebuildList := TreetoList(tree)
		require.Equal(t, tt.list, rebuildList)
	}
}

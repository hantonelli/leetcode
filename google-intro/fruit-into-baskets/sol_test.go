package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		tree  []int
		count int
	}{
		{
			tree:  []int{0, 1, 2, 2},
			count: 3,
		},
		{
			tree:  []int{1, 2, 3, 2, 2},
			count: 4,
		},
		{
			tree:  []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			count: 5,
		},
	}

	for _, te := range tests {
		require.Equal(t, te.count, totalFruit2(te.tree))
	}
}

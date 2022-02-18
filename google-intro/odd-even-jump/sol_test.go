package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		arr   []int
		count int
	}{
		{
			arr:   []int{10, 13, 12, 14, 15},
			count: 2,
		},
		{
			arr:   []int{2, 3, 1, 1, 4},
			count: 3,
		},
		{
			arr:   []int{5, 1, 3, 4, 2},
			count: 3,
		},
	}

	for _, te := range tests {
		require.Equal(t, te.count, oddEvenJumps(te.arr))
	}
}

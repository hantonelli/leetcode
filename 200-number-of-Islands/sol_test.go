package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	got := numIslands(grid)
	require.Equal(t, got, 1)
}

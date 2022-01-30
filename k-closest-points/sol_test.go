package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, [][]int{{-2, 2}}, KClosest2([][]int{{1, 3}, {-2, 2}}, 1))
}

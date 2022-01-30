package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanSum(t *testing.T) {
	require.Equal(t, true, CanSum(7, []int{2, 3}))
	require.Equal(t, true, CanSum(7, []int{5, 3, 4, 7}))
	require.Equal(t, false, CanSum(7, []int{2, 4}))
	require.Equal(t, true, CanSum(8, []int{2, 3, 5}))
	require.Equal(t, false, CanSum(300, []int{7, 14}))
}

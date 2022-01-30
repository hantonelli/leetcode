package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBestSum(t *testing.T) {
	require.Equal(t, []int{7}, BestSum(7, []int{5, 3, 4, 7}))
	require.Equal(t, []int{5, 3}, BestSum(8, []int{2, 3, 5}))
	require.Equal(t, []int{4, 4}, BestSum(8, []int{1, 4, 5}))
	require.Equal(t, []int{25, 25, 25, 25}, BestSum(100, []int{1, 2, 5, 25}))
}

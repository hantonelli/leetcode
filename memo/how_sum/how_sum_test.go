package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHowSum(t *testing.T) {
	var nilVal []int
	require.Equal(t, []int{3, 2, 2}, HowSum(7, []int{2, 3}))
	require.Equal(t, []int{4, 3}, HowSum(7, []int{5, 3, 4, 7}))
	require.Equal(t, nilVal, HowSum(7, []int{2, 4}))
	require.Equal(t, []int{2, 2, 2, 2}, HowSum(8, []int{2, 3, 5}))
	require.Equal(t, nilVal, HowSum(300, []int{7, 14}))
}

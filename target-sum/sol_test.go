package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, 5, []int{1, 1, 1, 1, 1})
	require.Equal(t, 1, []int{1})
}

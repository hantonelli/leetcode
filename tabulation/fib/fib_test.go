package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, 8, Fib(6))
	require.Equal(t, 13, Fib(7))
	require.Equal(t, 21, Fib(8))
	require.Equal(t, 12586269025, Fib(50))
}

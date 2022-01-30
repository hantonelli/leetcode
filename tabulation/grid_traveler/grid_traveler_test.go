package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGridTraveler(t *testing.T) {
	require.Equal(t, 1, GridTraveler(1, 1))
	require.Equal(t, 3, GridTraveler(2, 3))
	require.Equal(t, 3, GridTraveler(3, 2))
	require.Equal(t, 6, GridTraveler(3, 3))
	require.Equal(t, 2333606220, GridTraveler(18, 18))
}

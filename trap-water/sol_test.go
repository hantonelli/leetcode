package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			height:   []int{3, 1},
			expected: 0,
		},
		{
			height:   []int{3, 1, 3},
			expected: 2,
		},
		{
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, trap(tt.height))
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maxArea(tt.input))
	}
}

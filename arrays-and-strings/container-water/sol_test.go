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
			input:    []int{2, 3, 4, 5, 18, 17, 6},
			expected: 17,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maxArea(tt.input))
	}
}

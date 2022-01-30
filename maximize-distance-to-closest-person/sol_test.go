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
			input:    []int{1, 0, 0, 0, 1, 0, 1},
			expected: 2,
		},
		{
			input:    []int{1, 0, 0, 0},
			expected: 3,
		},
		{
			input:    []int{0, 0, 0, 1},
			expected: 3,
		},
		{
			input:    []int{1, 0, 0, 1},
			expected: 1,
		},
		{
			input:    []int{1, 0, 0, 0, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maxDistToClosest((tt.input)))
	}
}

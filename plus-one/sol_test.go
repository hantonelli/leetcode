package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{9, 9},
			expected: []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, plusOne(tt.input))
	}
}

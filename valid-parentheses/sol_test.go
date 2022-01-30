package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "()",
			expected: true,
		},
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
		{
			input:    "(",
			expected: false,
		},
		{
			input:    ")",
			expected: false,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, IsValid(tt.input))
	}
}

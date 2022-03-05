package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    " ",
			expected: 1,
		},
		{
			input:    "au",
			expected: 2,
		},
		{
			input:    "abba",
			expected: 2,
		},
		{
			input:    "au",
			expected: 2,
		},
		{
			input:    "nfpdmpi",
			expected: 5,
		},
		{
			input:    "tmmzuxt",
			expected: 5,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, lengthOfLongestSubstring(tt.input))
	}
}

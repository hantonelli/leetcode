package longest

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
			input:    "eceba",
			expected: 3,
		},
		{
			input:    "ccaabbb",
			expected: 5,
		},
		{
			input:    "ababacccccc",
			expected: 7,
		},
		{
			input:    "cdaba",
			expected: 3,
		},
		{
			input:    "abcbabac",
			expected: 4,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, lengthOfLongestSubstringTwoDistinctFast(tt.input))
	}
}

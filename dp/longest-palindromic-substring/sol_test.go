package longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "babad",
			expected: "bab",
		},
		{
			s:        "cbbd",
			expected: "bb",
		},
		{
			s:        "ab",
			expected: "a",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, longestPalindrome(tt.s))
	}
}

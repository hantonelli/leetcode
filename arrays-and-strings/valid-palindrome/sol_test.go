package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        " ",
			expected: true,
		},
		{
			s:        "a.",
			expected: true,
		},
		{
			s:        "ab",
			expected: false,
		},
		{
			s:        "0P",
			expected: false,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, isPalindrome(tt.s), tt.s)
	}
}

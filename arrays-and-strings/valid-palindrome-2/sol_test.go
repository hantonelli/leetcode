package valid_palindrome_2

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
			s:        "aba",
			expected: true,
		},
		{
			s:        "abca",
			expected: true,
		},
		{
			s:        "abc",
			expected: false,
		},
		{
			s: "eeccccbebaeeabebccceea",
		},
		{
			s:        "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			expected: true,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, validPalindromeMyFast(tt.s), tt.s)
	}
}

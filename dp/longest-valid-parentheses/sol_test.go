package longestvalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "(()",
			expected: 2,
		},
		{
			s:        ")()())",
			expected: 4,
		},
		{
			s:        "(()()(",
			expected: 4,
		},
		{
			s:        "(()())",
			expected: 6,
		},
		{
			s:        "",
			expected: 0,
		},
		{
			s:        "()(()",
			expected: 2,
		},
		{
			s:        "()()",
			expected: 4,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, longestValidParentheses(tt.s), tt.s)
	}
}

package string_to_integer

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
			s:        "42",
			expected: 42,
		},
		{
			s:        "   -42",
			expected: -42,
		},
		{
			s:        "--42",
			expected: 0,
		},
		{
			s:        "4-2",
			expected: 4,
		},
		{
			s:        "4193 with words",
			expected: 4193,
		},
		{
			s:        "20000000000000000000",
			expected: 2147483647,
		},
		{
			s:        "-20000000000000000000",
			expected: -2147483648,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, myAtoi(tt.s))
	}
}

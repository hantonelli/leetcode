package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"ab", "a"},
			expected: "a",
		},
		{
			strs:     []string{"abab", "aba", ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, longestCommonPrefix(tt.strs))
	}
}

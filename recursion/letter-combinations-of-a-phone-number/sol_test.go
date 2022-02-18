package lettercombinationsofaphonenumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		digits   string
		expected []string
	}{
		{
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits:   "",
			expected: []string{},
		},
		{
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, letterCombinations(tt.digits))
	}
}

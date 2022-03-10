package wordbreak

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, wordBreakMyFast(tt.s, tt.wordDict))
	}
}

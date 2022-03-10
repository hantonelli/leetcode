package wordladder2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  [][]string
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}},
		},
		{
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected:  [][]string{{"a", "c"}},
		},
		{
			beginWord: "teach",
			endWord:   "place",
			wordList:  []string{"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach"},
			expected:  [][]string{{"teach", "peach", "peace", "place"}},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, findLadders(tt.beginWord, tt.endWord, tt.wordList))
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			beginWord: "hot",
			endWord:   "dot",
			wordList:  []string{"hot", "dot", "dog"},
			expected:  2,
		},
		{
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"},
			expected:  3,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, ladderLength(tt.beginWord, tt.endWord, tt.wordList))
	}
}

func TestIsOneCharDiff(t *testing.T) {
	tests := []struct {
		x        string
		y        string
		expected bool
	}{
		{
			x:        "hit",
			y:        "cog",
			expected: false,
		},
		{
			x:        "hit",
			y:        "hot",
			expected: true,
		},
		{
			x:        "hit",
			y:        "cit",
			expected: true,
		},
		{
			x:        "hit",
			y:        "hiu",
			expected: true,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, isOneCharDiff(tt.x, tt.y))
	}

}

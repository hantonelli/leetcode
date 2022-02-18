package wordladder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var longWordList = []string{
	"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb",
	"kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br",
	"ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr",
	"po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh",
	"sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb",
	"ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz",
	"no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an",
	"me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi",
	"am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr",
	"pa", "he", "lr", "sq", "ye",
}

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
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"},
			expected:  3,
		},
		{
			beginWord: "qa",
			endWord:   "sq",
			wordList:  longWordList,
			expected:  5,
		},
		{
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dog"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, ladderLength2(tt.beginWord, tt.endWord, tt.wordList))
	}
}

func Benchmark_Works1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladderLength6("qa", "sq", longWordList)
	}
}

func Benchmark_Works2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladderLength5("qa", "sq", longWordList)
	}
}

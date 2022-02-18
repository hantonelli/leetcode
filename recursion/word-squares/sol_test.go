package wordsquares

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// b a l l
// a r e a
// l e a d
// l a d y

// w a l l

func Test(t *testing.T) {
	tests := []struct {
		words    []string
		expected [][]string
	}{
		{
			words: []string{"area", "lead", "wall", "lady", "ball"},
			expected: [][]string{
				{"wall", "area", "lead", "lady"},
				{"ball", "area", "lead", "lady"},
			},
		},
		{
			words:    []string{"abat", "baba", "atan", "atal"},
			expected: [][]string{{"baba", "abat", "baba", "atan"}, {"baba", "abat", "baba", "atal"}},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, wordSquares(tt.words))
	}
}

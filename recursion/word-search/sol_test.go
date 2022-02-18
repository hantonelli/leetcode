package wordsearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCCED",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "SEE",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCB",
			expected: false,
		},
		{
			board:    [][]byte{{'A', 'B'}},
			word:     "AB",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B'}},
			word:     "BA",
			expected: true,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, exist(tt.board, tt.word))
	}
}

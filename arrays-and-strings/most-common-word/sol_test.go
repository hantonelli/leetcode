package most_common_word

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		banned   []string
		expected string
	}{
		{
			s:        "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:   []string{"hit"},
			expected: "ball",
		},
		{
			s:        "Bob",
			banned:   []string{},
			expected: "bob",
		},
		{
			s:        "Bob Jon Jon",
			banned:   []string{},
			expected: "jon",
		},
		{
			s:        "Bob Jon Jon.",
			banned:   []string{},
			expected: "jon",
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, mostCommonWord(tt.s, tt.banned))
	}
}

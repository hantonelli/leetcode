package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, groupAnagrams(tt.strs))
	}
}

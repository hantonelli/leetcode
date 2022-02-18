package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		words    []string
		expected int
	}{
		{
			s:        "heeellooo",
			words:    []string{"hello", "hi", "helo"},
			expected: 1,
		},
		{
			s:        "zzzzzyyyyy",
			words:    []string{"zzyy", "zy", "zyy"},
			expected: 3,
		},
		{
			s:        "abcd",
			words:    []string{"abc"},
			expected: 0,
		},
		{
			s:        "dddiiiinnssssssoooo",
			words:    []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"},
			expected: 3,
		},
		{
			s:        "dddiiiinnssssssoooo",
			words:    []string{"dinnssoo", "ddiinnso", "ddiinnssoo"},
			expected: 3,
		},
		{
			s:        "aaa",
			words:    []string{"aaaa"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, expressiveWords(tt.s, tt.words))
	}
}

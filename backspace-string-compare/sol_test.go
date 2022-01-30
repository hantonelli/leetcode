package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "ab#c",
			t:        "ad#c",
			expected: true,
		},
		{
			s:        "ab##",
			t:        "c#d#",
			expected: true,
		},
		{
			s:        "ab##",
			t:        "c#d#",
			expected: true,
		},
		{
			s:        "a#b#",
			t:        "c##d#",
			expected: true,
		},
		{
			s:        "a#b#",
			t:        "#",
			expected: true,
		},
		{
			s:        "a",
			t:        "b",
			expected: false,
		},
		{
			s:        "a",
			t:        "",
			expected: false,
		},
		{
			s:        "",
			t:        "b",
			expected: false,
		},
		{
			s:        "bxj##tw",
			t:        "bxj###tw",
			expected: false,
		},
		{
			s:        "nzp#o#g",
			t:        "b#nzp#o#g",
			expected: true,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, backspaceCompare(tt.s, tt.t))
	}
}

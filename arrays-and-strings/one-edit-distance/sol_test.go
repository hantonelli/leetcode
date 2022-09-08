package one_edit_distance

import (
	"fmt"
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
			s:        "ab",
			t:        "acb",
			expected: true,
		},
		{
			s:        "abc",
			t:        "ab",
			expected: true,
		},
		{
			s:        "bc",
			t:        "b",
			expected: true,
		},
		{
			s:        "c",
			t:        "c",
			expected: false,
		},
		{
			s:        "abc",
			t:        "abc",
			expected: false,
		},
		{
			s:        "abc",
			t:        "bac",
			expected: false,
		},
		{
			s:        "abc",
			t:        "xbc",
			expected: true,
		},
		{
			s:        "abc",
			t:        "axc",
			expected: true,
		},
		{
			s:        "abc",
			t:        "abx",
			expected: true,
		},
		{
			s:        "",
			t:        "",
			expected: false,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, isOneEditDistance(tt.s, tt.t), fmt.Sprintf("s: %s, t: %s", tt.s, tt.t))
	}
}

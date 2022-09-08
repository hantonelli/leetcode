package add_binary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "10",
			b:        "1",
			expected: "11",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			a:        "11",
			b:        "11",
			expected: "110",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, addBinary(tt.a, tt.b))
	}
}

package crackingTheSafe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected string
	}{
		{
			n:        1,
			k:        2,
			expected: "01",
		},
		{
			n:        2,
			k:        2,
			expected: "00110",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, crackSafe(tt.n, tt.k))
	}
}

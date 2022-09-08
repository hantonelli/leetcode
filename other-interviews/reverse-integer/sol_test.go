package reverse_integer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{
			x:        123,
			expected: 321,
		},
		{
			x:        -123,
			expected: -321,
		},
		{
			x:        120,
			expected: 21,
		},
		{
			x:        2147483647,
			expected: 0,
		},
		{
			x:        -2147483648,
			expected: 0,
		},
		{
			x:        2000000003,
			expected: 0,
		},
		{
			x:        -2000000003,
			expected: 0,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, reverse(tt.x))
	}
}

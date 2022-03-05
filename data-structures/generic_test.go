package dataStructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        0,
			b:        1,
			expected: 1,
		},
		{
			a:        2,
			b:        1,
			expected: 2,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, Max(tt.a, tt.b))
	}
}

package nthtribonaccinumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{
			n:        4,
			expected: 4,
		},
		{
			n:        25,
			expected: 1389537,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, tribonacci2(tt.n))
	}
}

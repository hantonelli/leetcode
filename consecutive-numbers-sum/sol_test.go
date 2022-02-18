package consecutivenumberssum

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
			n:        1,
			expected: 1,
		},
		{
			n:        2,
			expected: 1,
		},
		{
			n:        5,
			expected: 2,
		},
		{
			n:        9,
			expected: 3,
		},
		{
			n:        10,
			expected: 2,
		},
		{
			n:        15,
			expected: 4,
		},
		{
			n:        57265,
			expected: 8,
		},
		{
			n:        855877922,
			expected: 4,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, consecutiveNumbersSum(tt.n))
	}
}

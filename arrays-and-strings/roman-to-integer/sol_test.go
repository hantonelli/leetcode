package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "III",
			expected: 3,
		},
		{
			s:        "LVIII",
			expected: 58,
		},
		{
			s:        "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, romanToInt(tt.s))
	}
}

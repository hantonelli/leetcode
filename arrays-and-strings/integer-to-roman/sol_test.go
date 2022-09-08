package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{
			num:      3,
			expected: "III",
		},
		{
			num:      58,
			expected: "LVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, intToRoman(tt.num))
	}
}

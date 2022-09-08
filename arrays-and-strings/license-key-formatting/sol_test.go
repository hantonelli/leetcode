package license_key_formatting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected string
	}{
		{
			s:        "5F3Z-2e-9-w",
			k:        4,
			expected: "5F3Z-2E9W",
		},
		{
			s:        "2-5g-3-J",
			k:        2,
			expected: "2-5G-3J",
		},
		{
			s:        "-----2----5g-3-J----",
			k:        2,
			expected: "2-5G-3J",
		},
		{
			s:        "r",
			k:        1,
			expected: "R",
		},
		{
			s:        "aaaa",
			k:        2,
			expected: "AA-AA",
		},
		{
			s:        "---",
			k:        3,
			expected: "",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, licenseKeyFormatting(tt.s, tt.k))
	}
}

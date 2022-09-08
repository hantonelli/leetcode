package validate_ip_address

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		queryIP  string
		expected string
	}{
		{
			queryIP:  "192.168.1.0",
			expected: IPV4,
		},
		{
			queryIP:  "2001:0db8:85a3:0:0:8A2E:0370:7334",
			expected: IPV6,
		},
		{
			queryIP:  "256.256.256.256",
			expected: NEITHER,
		},
		{
			queryIP:  "192.168.01.1",
			expected: NEITHER,
		},
		{
			queryIP:  "192.168.1.00",
			expected: NEITHER,
		},
		{
			queryIP:  "192.168@1.1",
			expected: NEITHER,
		},
		{
			queryIP:  "2001:0db8:85a3::8A2E:037j:7334",
			expected: NEITHER,
		},
		{
			queryIP:  "02001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: NEITHER,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, validIPAddress(tt.queryIP), tt.queryIP)
	}
}

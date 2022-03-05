package findreplacestring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s       string
		indices []int
		sources []string
		targets []string
		expect  string
	}{
		{
			s:       "abcd",
			indices: []int{0, 2},
			sources: []string{"a", "cd"},
			targets: []string{"eee", "ffff"},
			expect:  "eeebffff",
		},
		{
			s:       "abcd",
			indices: []int{0, 2},
			sources: []string{"ab", "ec"},
			targets: []string{"eee", "ffff"},
			expect:  "eeecd",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, findReplaceString(tt.s, tt.indices, tt.sources, tt.targets))
	}
}

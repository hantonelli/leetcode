package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect string
	}{
		{
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
		{
			s:      "a",
			t:      "a",
			expect: "a",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, minWindow(tt.s, tt.t))
	}
}

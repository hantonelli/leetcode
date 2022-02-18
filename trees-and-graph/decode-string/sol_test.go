package decodeString

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			input:    "10[a]",
			expected: "aaaaaaaaaa",
		},
		{
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, decodeString(tt.input))
	}
}

func Benchmark_Sol(b *testing.B) {
	input := "2[abc]3[cd]ef2[abc]3[cd]ef2[abc]3[cd]ef"
	for i := 0; i < b.N; i++ {
		decodeString(input)
	}
}

func Benchmark_SolFast(b *testing.B) {
	input := "2[abc]3[cd]ef2[abc]3[cd]ef2[abc]3[cd]ef"
	for i := 0; i < b.N; i++ {
		decodeString3(input)
	}
}

func Benchmark_Fastest(b *testing.B) {
	input := "2[abc]3[cd]ef2[abc]3[cd]ef2[abc]3[cd]ef"
	for i := 0; i < b.N; i++ {
		decodeString2(input)
	}
}

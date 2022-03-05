package longestvalidparentheses

import "testing"

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParentheses("(()())")
	}
}

func Benchmark_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestValidParenthesesFast("(()())")
	}
}

package string_to_integer

import "testing"

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("4193 with words")
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoiFast("4193 with words")
	}
}

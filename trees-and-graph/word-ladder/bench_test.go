package wordladder

import "testing"

func Benchmark_Works1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladderLength6("qa", "sq", longWordList)
	}
}

func Benchmark_Works2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ladderLength5("qa", "sq", longWordList)
	}
}

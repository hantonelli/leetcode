package wordladder2

import "testing"

func Benchmark_Works1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findLadders("teach", "place", []string{"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach"})
	}
}

func Benchmark_Works2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findLaddersFast("teach", "place", []string{"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach"})
	}
}

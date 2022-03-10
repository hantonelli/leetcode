package crackingTheSafe

import "testing"

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crackSafe(4, 10)
	}
}
func Benchmark_SolG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crackSafeFast(4, 10)
	}
}

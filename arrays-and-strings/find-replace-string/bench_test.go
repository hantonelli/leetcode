package findreplacestring

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"})
	}
}
func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findReplaceStringFast("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"})
	}
}
func Benchmark_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findReplaceStringSmall("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"})
	}
}

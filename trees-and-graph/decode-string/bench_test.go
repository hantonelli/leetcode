package decodeString

import "testing"

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

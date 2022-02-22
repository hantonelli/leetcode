package wordsearch2

import "testing"

func Benchmark_Sol(b *testing.B) {
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}
	for i := 0; i < b.N; i++ {
		findWords(board, words)
	}
}

func Benchmark_Fast(b *testing.B) {
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}
	for i := 0; i < b.N; i++ {
		findWords2(board, words)
	}
}

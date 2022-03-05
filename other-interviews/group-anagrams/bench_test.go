package groupanagrams

import (
	"testing"
)

var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagramsMyFast(strs)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagramsFast(strs)
	}
}

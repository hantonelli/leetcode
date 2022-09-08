package one_edit_distance

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var (
	pos = 500
	str = ds.RandomWord(1000)
)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str2 := str[:pos] + str[pos+1:]
		isOneEditDistance(str, str2)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str2 := str[:pos] + str[pos+1:]
		isOneEditDistanceFast(str, str2)
	}
}

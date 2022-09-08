package permutations

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(10)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permute(base)
	}
}

func Benchmark_BT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permuteBT(base)
	}
}

func Benchmark_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permuteSmall(base)
	}
}

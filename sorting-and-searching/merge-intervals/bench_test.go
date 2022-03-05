package mergeintervals

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base [][]int = ds.RandArrOfArr(10000, 100)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge(base)
	}
}

func Benchmark_SolFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeFast(base)
	}
}

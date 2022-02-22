package mergeintervals

import (
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
)

var base [][]int = dataStructures.RandArrOfArr(10000, 100)

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

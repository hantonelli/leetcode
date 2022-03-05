package maximumsubarray

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(1000000)

func Benchmark_Works1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubArray(base)
	}
}

func Benchmark_Works2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubArrayFast(base)
	}
}

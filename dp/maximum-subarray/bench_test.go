package maximumsubarray

import (
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
)

var base []int = dataStructures.RandArr(1000000)

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

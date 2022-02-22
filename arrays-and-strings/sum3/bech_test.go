package sum3

import (
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
)

var base []int = dataStructures.RandArr(100)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum(base)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSumFast(base)
	}
}

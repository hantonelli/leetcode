package two_sum

import "testing"

func Benchmark_Sol(b *testing.B) {
	input := []int{0, 11, 15, 2, 11, 66, 33, 7}
	for i := 0; i < b.N; i++ {
		twoSum(input, 9)
	}
}

func Benchmark_Fast(b *testing.B) {
	input := []int{0, 11, 15, 2, 11, 66, 33, 7}
	for i := 0; i < b.N; i++ {
		twoSum2(input, 9)
	}
}

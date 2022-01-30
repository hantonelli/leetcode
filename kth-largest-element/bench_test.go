package main

import (
	"math/rand"
	"testing"
)

func getInput(n int) []int {
	input := make([]int, n)
	for i := 0; i < n; i++ {
		input[i] = rand.Intn(1000)
	}
	return input
}

var input = getInput(500000)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findKthLargest(input, 1000)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findKthLargest2(input, 1000)
	}
}

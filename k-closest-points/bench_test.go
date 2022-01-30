package main

import (
	"math/rand"
	"testing"
)

func getInput(n int) [][]int {
	input := make([][]int, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		input[i] = []int{x, y}
	}
	return input
}

var input = getInput(500000)

// func Benchmark_Fast(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		KClosestFast(input, 100)
// 	}
// }

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KClosest(input, 100)
	}
}

func Benchmark_Sol3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KClosest4(input, 100)
	}
}

func Benchmark_Fast2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KClosestFast2(input, 100)
	}
}

func Benchmark_Fast3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KClosestFast3(input, 100)
	}
}

func Benchmark_MineFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KClosest2(input, 100)
	}
}

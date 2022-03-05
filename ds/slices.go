package ds

import (
	"math"
	"math/rand"
)

func RandArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(math.MaxInt64)
	}
	return arr
}

func RandArrOfArr(size, internal int) [][]int {
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, internal)
		for j := 0; j < internal; j++ {
			arr[i][j] = rand.Intn(math.MaxInt64)
		}
	}
	return arr
}

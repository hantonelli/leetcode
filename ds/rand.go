package ds

import (
	"math"
	"math/rand"
)

var (
	letters       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	lettersULCase = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLCase  = []rune("abcdefghijklmnopqrstuvwxyz")
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

func RandomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomWord(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = lettersULCase[rand.Intn(len(lettersULCase))]
	}
	return string(s)
}

func RandomWordLowercase(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = lettersLCase[rand.Intn(len(lettersLCase))]
	}
	return string(s)
}

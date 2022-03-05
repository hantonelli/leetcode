package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

type Ordered interface {
	constraints.Integer | constraints.Float | ~string | rune | byte
}

func Quicksort[T Ordered](a []T) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	Quicksort(a[:left])
	Quicksort(a[left+1:])
}

func QuicksortFn[T any](a []T, cmp func(i, j T) bool) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if cmp(a[i], a[right]) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	QuicksortFn(a[:left], cmp)
	QuicksortFn(a[left+1:], cmp)
}

package minimumcosttohirekworkers

import (
	"container/heap"
	"math"
	"sort"
)

type intHeap struct {
	sum int
	arr []int
}

func (h intHeap) Len() int           { return len(h.arr) }
func (h intHeap) Less(i, j int) bool { return h.arr[i] > h.arr[j] }
func (h intHeap) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.sum += x.(int)
	h.arr = append(h.arr, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.sum -= x
	h.arr = old[0 : n-1]
	return x
}

func (h *intHeap) Sum() int {
	return h.sum
}

func (h *intHeap) Size() int {
	return len(h.arr)
}

func (h *intHeap) Peek() int {
	return h.arr[0]
}

type worker struct {
	quality        int
	costPerQuality float64
}

func mincostToHireWorkers3(quality []int, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := 0; i < len(quality); i++ {
		workers[i] = worker{
			quality:        quality[i],
			costPerQuality: float64(wage[i]) / float64(quality[i]),
		}
	}
	sort.SliceStable(workers, func(i, j int) bool { return workers[i].costPerQuality < workers[j].costPerQuality })

	var minCost float64 = -1
	h := &intHeap{}
	heap.Init(h)
	for i := 0; i < len(workers); i++ {
		if h.Size() < k {
			heap.Push(h, workers[i].quality)
			if h.Size() == k {
				currCost := float64(h.Sum()) * workers[i].costPerQuality
				if minCost == -1 || currCost < minCost {
					minCost = currCost
				}
			}
		} else {
			if workers[i].quality < h.Peek() {
				heap.Pop(h)
				heap.Push(h, workers[i].quality)

				currCost := float64(h.Sum()) * workers[i].costPerQuality
				if minCost == -1 || currCost < minCost {
					minCost = currCost
				}
			}
		}
	}
	return math.Round(minCost*100000) / 100000
}

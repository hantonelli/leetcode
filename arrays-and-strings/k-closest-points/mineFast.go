package main

import (
	"container/heap"
	"math"
)

type Point struct {
	x    int
	y    int
	dist float64
}

type DistHeap []Point

func (h DistHeap) Len() int           { return len(h) }
func (h DistHeap) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h DistHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DistHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *DistHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KClosest2(points [][]int, k int) [][]int {
	if len(points) < k {
		return points
	}
	var dist float64
	res := make(DistHeap, k)
	for i := 0; i < k; i++ {
		dist = math.Sqrt(float64(points[i][0]*points[i][0] + points[i][1]*points[i][1]))
		res[i] = Point{x: points[i][0], y: points[i][1], dist: dist}
	}
	heap.Init(&res)

	for j := k; j < len(points); j++ {
		dist = math.Sqrt(float64(points[j][0]*points[j][0] + points[j][1]*points[j][1]))

		if len(res) < k || dist < res[k-1].dist {
			if len(res) == k {
				heap.Pop(&res)
			}
			heap.Push(&res, Point{x: points[j][0], y: points[j][1], dist: dist})
		}
	}
	res2 := make([][]int, k)
	for l := 0; l < k; l++ {
		res2[l] = []int{res[l].x, res[l].y}
	}
	return res2
}

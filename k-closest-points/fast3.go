package main

import "math"

func KClosestFast3(points [][]int, k int) [][]int {
	heap := maxHeap3{
		array: make([]point3, len(points)+1),
	}
	for i := range points {
		x := points[i][0]
		y := points[i][1]
		d := math.Sqrt(float64(x*x + y*y))
		p := point3{
			x: x,
			y: y,
			d: d,
		}
		if heap.size < k {
			heap.push(p)
		} else if heap.max() > d {
			heap.pop()
			heap.push(p)
		}
	}
	closest := make([][]int, 0, heap.size)
	for heap.size > 0 {
		p := heap.pop()
		closest = append(closest, []int{p.x, p.y})
	}
	return closest
}

type point3 struct {
	x int
	y int
	d float64
}

type maxHeap3 struct {
	size  int
	array []point3
}

func (h *maxHeap3) max() float64 {
	return h.array[1].d
}

func (h *maxHeap3) push(p point3) {
	h.size++
	h.array[h.size] = p
	current := h.size
	parent := current / 2
	for parent > 0 {
		if h.array[current].d <= h.array[parent].d {
			break
		}
		tmp := h.array[current]
		h.array[current] = h.array[parent]
		h.array[parent] = tmp
		current = parent
		parent = current / 2
	}
}

func (h *maxHeap3) pop() point3 {
	result := h.array[1]
	h.array[1] = h.array[h.size]
	h.size--
	current := 1
	minIndex := current * 2
	for minIndex <= h.size {
		if h.size > minIndex && h.array[minIndex].d < h.array[minIndex+1].d {
			minIndex++
		}
		if h.array[minIndex].d <= h.array[current].d {
			break
		}
		tmp := h.array[current]
		h.array[current] = h.array[minIndex]
		h.array[minIndex] = tmp
		current = minIndex
		minIndex = current * 2
	}
	return result
}

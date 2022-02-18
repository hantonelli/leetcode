package main

import "math/rand"

func KClosestFast(points [][]int, k int) [][]int {
	if k == len(points) {
		return points
	}

	dist := make([]int, len(points))

	for i := 0; i < len(points); i++ {
		dist[i] = points[i][0]*points[i][0] + points[i][1]*points[i][1]
	}
	partitionbyk(dist, points, k)
	return points[:k]
}

func partitionbyk(dist []int, points [][]int, k int) {
	r := rand.Intn(len(dist))
	dist[0], dist[r] = dist[r], dist[0]
	points[0], points[r] = points[r], points[0]
	i := 1
	l := len(dist) - 1
	for ; i <= l; i++ {
		if dist[i] > dist[0] {
			dist[i], dist[l] = dist[l], dist[i]
			points[i], points[l] = points[l], points[i]
			l--
			i--
		}
	}
	l++
	if l == k {
		return
	} else if l > k {
		partitionbyk(dist[:l], points[:l], k)
	} else {
		partitionbyk(dist[l:], points[l:], k-l)

	}
}

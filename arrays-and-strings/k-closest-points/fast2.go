package main

import "sort"

func KClosestFast2(points [][]int, k int) [][]int { // Time: O(nlogn), Space: O(n)
	if len(points) == 0 {
		return points
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:k]
}

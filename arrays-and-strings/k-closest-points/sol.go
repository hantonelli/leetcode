package main

import (
	"math"
	"sort"
)

type item struct {
	distance float64
	pointX   int
	pointY   int
}

type items []item

func (a items) Len() int           { return len(a) }
func (a items) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a items) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func KClosest(points [][]int, k int) [][]int {
	if len(points) < k {
		return points
	}
	var dist float64
	res := make([]item, k)
	for i := 0; i < k; i++ {
		dist = math.Sqrt(float64(points[i][0])*float64(points[i][0]) + float64(points[i][1])*float64(points[i][1]))
		res[i] = item{pointX: points[i][0], pointY: points[i][1], distance: dist}
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].distance < res[j].distance
	})
	for j := k; j < len(points); j++ {
		dist = math.Sqrt(float64(points[j][0])*float64(points[j][0]) + float64(points[j][1])*float64(points[j][1]))
		if dist < res[k-1].distance {
			res[k-1] = item{pointX: points[j][0], pointY: points[j][1], distance: dist}

			sort.SliceStable(res, func(i, j int) bool {
				return res[i].distance < res[j].distance
			})
		}
	}
	res2 := make([][]int, k)
	for l := 0; l < k; l++ {
		res2[l] = []int{res[l].pointX, res[l].pointY}
	}
	return res2
}

func KClosest4(points [][]int, k int) [][]int {
	if len(points) < k {
		return points
	}
	var dist float64
	res := make(items, k)
	for i := 0; i < k; i++ {
		dist = math.Sqrt(float64(points[i][0])*float64(points[i][0]) + float64(points[i][1])*float64(points[i][1]))
		res[i] = item{pointX: points[i][0], pointY: points[i][1], distance: dist}
	}
	sort.Sort(res)
	for j := k; j < len(points); j++ {
		dist = math.Sqrt(float64(points[j][0])*float64(points[j][0]) + float64(points[j][1])*float64(points[j][1]))
		if dist < res[k-1].distance {
			res[k-1] = item{pointX: points[j][0], pointY: points[j][1], distance: dist}
			sort.Sort(res)
		}
	}
	res2 := make([][]int, k)
	for l := 0; l < k; l++ {
		res2[l] = []int{res[l].pointX, res[l].pointY}
	}
	return res2
}

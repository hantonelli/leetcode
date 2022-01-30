package main

// import (
// 	"fmt"
// 	"sort"
// )

// func canPartition(numbs []int) bool {
// 	total := 0
// 	for _, num := range numbs {
// 		total += num
// 	}
// 	if total%2 != 0 {
// 		return false
// 	}
// 	target := total / 2

// 	sort.Ints(numbs)

// 	memo := map[string]bool{}
// 	return canSum(target, numbs, memo)
// }

// func canSum(target int, numbs []int, memo map[string]bool) bool {
// 	if v, ok := memo[fmt.Sprintf("%v;%v", target, numbs)]; ok {
// 		return v
// 	}
// 	if target == 0 {
// 		return true
// 	}
// 	if target < 0 {
// 		return false
// 	}
// 	for i := 0; i < len(numbs); i++ {
// 		subNums := remove(numbs, i)
// 		if canSum(target-numbs[i], subNums, memo) {
// 			memo[fmt.Sprintf("%v;%v", target, numbs)] = true
// 			return true
// 		}
// 	}
// 	memo[fmt.Sprintf("%v;%v", target, numbs)] = false
// 	return false
// }

// func remove(src []int, i int) []int {
// 	dest := make([]int, 0, len(src)-1)
// 	dest = append(dest, src[:i]...)
// 	return append(dest, src[i+1:]...)
// }

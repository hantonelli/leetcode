package main

import (
	"fmt"
	"sort"
)

func CanPartition2(numbs []int) bool {
	total := 0
	for _, num := range numbs {
		total += num
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2

	sort.Ints(numbs)

	memo := map[string]bool{}
	return canSum2(target, 0, 0, numbs, memo)
}

func canSum2(targetSum, currIdx, currSum int, numbs []int, memo map[string]bool) bool {
	if targetSum == currSum {
		return true
	}
	if len(numbs) <= currIdx+1 {
		return false
	}

	key1 := fmt.Sprintf("%d;%d", currIdx+1, currSum)
	canSumNotUsing, ok := memo[key1]
	if !ok {
		memo[key1] = canSum2(currIdx+1, currSum, targetSum, numbs, memo)
	}

	key2 := fmt.Sprintf("%d;%d", currIdx+1, currSum-numbs[currIdx])
	canSumUsing, ok2 := memo[key2]
	if !ok2 {
		memo[key2] = canSum2(currIdx+1, currSum-numbs[currIdx], targetSum, numbs, memo)
	}

	return canSumNotUsing || canSumUsing
}

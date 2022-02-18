package deleteandearn

import "sort"

func deleteAndEarn(nums []int) int {
	numsMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] += nums[i]
	}
	numKeys := make([]int, 0, len(numsMap))
	for k := range numsMap {
		numKeys = append(numKeys, k)
	}
	sort.Ints(numKeys)
	memo := make(map[int]int, len(numKeys))
	return dp(0, numKeys, numsMap, memo)
}

func dp(i int, numKeys []int, numsMap map[int]int, memo map[int]int) int {
	if len(numKeys) <= i {
		return 0
	}
	if v, ok := memo[i]; ok {
		return v
	}
	iWin := numsMap[numKeys[i]]
	var win int
	if i+1 < len(numKeys) && numKeys[i+1] == numKeys[i]+1 {
		win = max(dp(i+1, numKeys, numsMap, memo), iWin+dp(i+2, numKeys, numsMap, memo))
	} else {
		win = iWin + dp(i+1, numKeys, numsMap, memo)
	}
	memo[i] = win
	return win
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

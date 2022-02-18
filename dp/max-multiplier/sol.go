package maxmultiplier

import "fmt"

func maximumScore(nums []int, multipliers []int) int {
	memo := make(map[string]int, len(multipliers))
	return dp(0, len(nums)-1, 0, nums, multipliers, memo)
}

func dp(nStart, nEnd, mI int, n []int, m []int, memo map[string]int) int {
	if len(m) <= mI {
		return 0
	}
	key := fmt.Sprintf("%d-%d", nStart, nEnd)
	if v, ok := memo[key]; ok {
		return v
	}

	dpStart, ok2 := memo[fmt.Sprintf("%d-%d", nStart+1, nEnd)]
	if !ok2 {
		dpStart = dp(nStart+1, nEnd, mI+1, n, m, memo)
	}
	resStart := n[nStart]*m[mI] + dpStart

	dpEnd, ok3 := memo[fmt.Sprintf("%d-%d", nStart, nEnd-1)]
	if !ok3 {
		dpEnd = dp(nStart, nEnd-1, mI+1, n, m, memo)
	}
	resEnd := n[nEnd]*m[mI] + dpEnd

	res := max(resStart, resEnd)
	memo[key] = res
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

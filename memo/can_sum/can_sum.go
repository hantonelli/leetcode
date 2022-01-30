package main

func CanSum(targetSum int, numbers []int) bool {
	memo := map[int]bool{}
	return canSum(targetSum, numbers, memo)
}

func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for i := 0; i < len(numbers); i++ {
		if canSum(targetSum-numbers[i], numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}

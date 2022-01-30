package main

func HowSum(targetSum int, numbers []int) []int {
	memo := map[int][]int{}
	return howSum(targetSum, numbers, memo)
}

func howSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for i := 0; i < len(numbers); i++ {
		arrRest := howSum(targetSum-numbers[i], numbers, memo)
		if arrRest != nil {
			newArr := append(arrRest, numbers[i])
			memo[targetSum] = newArr
			return memo[targetSum]
		}
	}

	memo[targetSum] = nil
	return nil
}

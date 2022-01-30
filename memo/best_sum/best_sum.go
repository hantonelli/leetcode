package main

func BestSum(target int, numbers []int) []int {
	memo := map[int][]int{}
	return bestSum(target, numbers, memo)
}

func bestSum(target int, numbers []int, memo map[int][]int) []int {
	if v, ok := memo[target]; ok {
		return v
	}
	if target < 0 {
		return nil
	}
	if target == 0 {
		return []int{}
	}
	var bestArr []int
	for _, num := range numbers {
		if arrSumRest := bestSum(target-num, numbers, memo); arrSumRest != nil {
			newArr := append(arrSumRest, num)
			if bestArr == nil || len(newArr) < len(bestArr) {
				bestArr = newArr
			}
		}
	}
	memo[target] = bestArr
	return bestArr
}

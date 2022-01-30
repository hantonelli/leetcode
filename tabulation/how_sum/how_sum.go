package main

func HowSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	values := make([][]int, targetSum+1)
	values[0] = []int{}

	for i := 0; i < targetSum+1; i++ {
		if values[i] != nil {
			for _, num := range numbers {
				if i+num < targetSum+1 {
					values[i+num] = append(values[i], num)
				}
			}
		}
	}

	return values[targetSum]
}

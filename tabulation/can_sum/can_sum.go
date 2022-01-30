package main

func CanSum(targetSum int, numbers []int) bool {
	values := make([]bool, targetSum+1)
	values[0] = true

	for i := 0; i < targetSum+1; i++ {
		if values[i] {
			for _, num := range numbers {
				if i+num < targetSum+1 {
					values[i+num] = true
				}
			}
		}
	}

	return values[targetSum]
}

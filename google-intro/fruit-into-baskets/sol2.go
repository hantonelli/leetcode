package main

func totalFruit2(fruits []int) int {
	var max int
	for i := 0; i < len(fruits); i++ {
		type1 := fruits[i]
		currCount := 1
		type2 := -1
		for j := i + 1; j < len(fruits); j++ {
			if type2 == -1 && type1 != fruits[j] {
				type2 = fruits[j]
			}
			if fruits[j] == type1 || fruits[j] == type2 {
				currCount++
			} else {
				break
			}
		}
		if max < currCount {
			max = currCount
		}
	}
	return max
}

package main

func totalFruit(fruits []int) int {
	otherType, tailType := -1, -1
	tailCount, currCount := 0, 0
	result := 0

	for _, currType := range fruits {
		currCount++
		if currType != tailType && currType != otherType {
			currCount = tailCount + 1
		}

		if result < currCount {
			result = currCount
		}

		tailCount++
		if currType != tailType {
			otherType, tailType = tailType, currType
			tailCount = 1
		}
	}

	return result
}

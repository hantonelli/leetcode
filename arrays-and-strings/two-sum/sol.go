package two_sum

func twoSum(numbers []int, target int) []int {
	// sort.Ints(numbers) // It's sorted already
	if len(numbers) < 2 {
		return []int{}
	}
	res := []int{}
	left, right := 0, len(numbers)-1
	for left < right {
		diff := target - numbers[left] - numbers[right]
		if 0 == diff {
			return []int{left + 1, right + 1}
		}
		if 0 < diff {
			left++
		} else {
			right--
		}
	}
	return res
}

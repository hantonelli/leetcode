package missing_number

func missingNumberFast(nums []int) int {
	var diff int
	for i, num := range nums {
		diff += num - i
	}
	return len(nums) - diff
}

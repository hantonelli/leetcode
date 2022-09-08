package two_sum

func twoSumMap(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	rest := target - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == rest {
			return []int{0, i}
		}
		numMap[nums[i]] = i
	}
	for i := 1; i < len(nums); i++ {
		rest = target - nums[i]
		if v, ok := numMap[rest]; ok && v != i {
			return []int{i, numMap[rest]}
		}
	}
	return nil
}

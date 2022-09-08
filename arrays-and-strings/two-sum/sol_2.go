package two_sum

func twoSum2(nums []int, target int) []int {
	positionMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := positionMap[diff]; ok {
			return []int{i, v}
		}
		for j := i + 1; j < len(nums); j++ {
			if diff-nums[j] == 0 {
				return []int{i, j}
			}
			positionMap[nums[j]] = j
		}
	}
	return nil
}

package move_zeroes

func moveZeroesFast(nums []int) {
	lastzero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastzero] = nums[i]
			lastzero++
		}
	}
	for i := lastzero; i < len(nums); i++ {
		nums[i] = 0
	}
}

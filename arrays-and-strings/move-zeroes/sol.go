package move_zeroes

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			restZeros := true
			for j := i; j+1 < len(nums); j++ {
				if nums[j+1] != 0 {
					restZeros = false
				}
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			if restZeros {
				return
			}
			i--
		}
	}
}

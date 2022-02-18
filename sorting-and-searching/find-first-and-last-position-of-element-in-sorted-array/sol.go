package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	var i int
	var finishNext bool
	right, left := 0, len(nums)

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	for {
		i = int((right + left) / 2)
		if nums[i] == target {
			resStart, resFinish := i, i
			for 0 <= resStart-1 && nums[resStart-1] == target {
				resStart--
			}
			for resFinish+1 < len(nums) && nums[resFinish+1] == target {
				resFinish++
			}
			return []int{resStart, resFinish}
		}
		if nums[i] < target {
			right = i
		} else {
			left = i
		}
		if finishNext {
			break
		}
		if left-1 <= right {
			finishNext = true
		}
	}
	return []int{-1, -1}
}

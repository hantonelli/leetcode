package nextpermutation

func nextPermutation(nums []int) {
	// Find the first decreasing number iterating from right to left.
	fd := len(nums) - 2
	for 0 <= fd && nums[fd+1] <= nums[fd] {
		fd--
	}
	// If the fd is 0, then all the numbers are already in decreasing order
	// In the case the full number needs to be reversed.
	if 0 <= fd {
		// If there is an intermediate decreading number, then we need to find the
		// smallest number that is bigger than fd, iterating from right to left
		j := len(nums) - 1
		for nums[j] <= nums[fd] {
			j--
		}
		nums[fd], nums[j] = nums[j], nums[fd]
	}

	// Reverse
	fd, j := fd+1, len(nums)-1
	for fd < j {
		nums[fd], nums[j] = nums[j], nums[fd]
		fd++
		j--
	}
}

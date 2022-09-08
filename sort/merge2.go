package sort

// func MergeSort2(nums []int) []int {
// 	n := len(nums)
// 	result := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		result[i] = i
// 	}
// 	if len(nums) <= 1 {
// 		return nums
// 	}
// 	middle := len(nums) / 2
// 	mergeSort2(0, middle, nums, result)
// 	mergeSort2(middle, len(nums), nums, result)
// 	merge2(0, len(nums), middle, nums, result)
// 	return result
// }

// func mergeSort2(left, right int, nums, result []int) {
// 	if right-left <= 1 {
// 		return
// 	}
// 	mid := (left + right) / 2
// 	mergeSort2(left, mid, nums, result)
// 	mergeSort2(mid, right, nums, result)
// 	merge2(left, right, mid, nums, result)
// }

// func merge2(left, right, mid int, nums, result []int) {
// 	// merge [left, mid) and [mid, right)
// 	i := left // current index for the left array
// 	j := mid  // current index for the right array
// 	// use temp to temporarily store sorted array
// 	pos := 0
// 	for i < mid && j < right {
// 		if nums[i] <= nums[j] {
// 			result[left+pos] = nums[i]
// 			pos++
// 			i++
// 		} else {
// 			result[left+pos] = nums[i]
// 			pos++
// 			j++
// 		}
// 	}
// 	for i < mid {
// 		// j - mid numbers jump to the left side of indices[i]
// 		result[left+pos] = result[i]
// 		pos++
// 		i++
// 	}
// 	for j < right {
// 		result[left+pos] = result[j]
// 		pos++
// 		j++
// 	}
// 	// restore from temp
// 	for k := left; k < right; k++ {
// 		result[left+pos] = nums[k-left]
// 		pos++
// 	}
// }

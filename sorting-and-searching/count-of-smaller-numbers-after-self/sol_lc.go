package countofsmallernumbersafterself

func countSmallerLC(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	indices := make([]int, n) // record the index. we are going to sort this array
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	// sort indices with their corresponding values in nums, i.e., nums[indices[i]]
	mergeSortLC(indices, 0, n, result, nums)
	// change int[] to List to return
	resultToReturn := make([]int, n)
	for i := range result {
		resultToReturn[i] = result[i]
	}
	return resultToReturn
}

func mergeSortLC(indices []int, left, right int, result, nums []int) {
	if right-left <= 1 {
		return
	}
	mid := (left + right) / 2
	mergeSortLC(indices, left, mid, result, nums)
	mergeSortLC(indices, mid, right, result, nums)
	mergeLC(indices, left, right, mid, result, nums)
}

func mergeLC(indices []int, left, right, mid int, result, nums []int) {
	// merge [left, mid) and [mid, right)
	i := left // current index for the left array
	j := mid  // current index for the right array
	// use temp to temporarily store sorted array
	pos := 0
	temp := make([]int, right-left)
	for i < mid && j < right {
		if nums[indices[i]] <= nums[indices[j]] {
			// j - mid numbers jump to the left side of indices[i]
			result[indices[i]] += j - mid
			temp[pos] = indices[i]
			pos++
			i++
		} else {
			temp[pos] = indices[j]
			pos++
			j++
		}
	}
	// when one of the subarrays is empty
	for i < mid {
		// j - mid numbers jump to the left side of indices[i]
		result[indices[i]] += j - mid
		temp[pos] = indices[i]
		pos++
		i++
	}
	for j < right {
		temp[pos] = indices[j]
		pos++
		j++
	}
	// restore from temp
	for k := left; k < right; k++ {
		indices[k] = temp[k-left]
	}
}

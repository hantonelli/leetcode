package main

import "sort"

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		sort.Ints(nums)
		return nums[0]
	}
	kList := make([]int, k)
	for i := 0; i < k; i++ {
		kList[i] = nums[i]
	}
	sort.Ints(kList)

	for j := k; j < len(nums); j++ {
		if kList[0] < nums[j] {
			kList[0] = nums[j]
			shortSort(kList)
		}
	}
	return kList[0]
}

func shortSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}

package permutations

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	// other := make([]int, len(nums)-1)
	// res := [][]int{}
	res := make([][]int, 0, factorial(len(nums)))
	for i := 0; i < len(nums); i++ {
		// for j := 0; j < i; j++ {
		// 	other[j] = nums[j]
		// }
		// for j := i + 1; j < len(nums); j++ {
		// 	other[j-1] = nums[j]
		// }
		other := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		opts := permute(other)
		for _, opt := range opts {
			res = append(res, append(opt, nums[i]))
		}
	}
	return res
}

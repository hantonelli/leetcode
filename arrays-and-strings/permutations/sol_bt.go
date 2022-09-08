package permutations

// func factorial(n int) int {
// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}
// 	return result
// }

func permuteBT(nums []int) [][]int {
	res := make([][]int, 0, factorial(len(nums)))
	backtrackBT(nums, make([]int, 0, len(nums)), make([]bool, len(nums)), &res)
	return res
}

func backtrackBT(nums, running []int, used []bool, res *[][]int) {
	if len(running) == len(nums) {
		buf := make([]int, len(running))
		copy(buf, running)
		*res = append(*res, buf)
		return
	}

	for i := range nums {
		if !used[i] {
			used[i] = true
			backtrackBT(nums, append(running, nums[i]), used, res)
			used[i] = false
		}
	}
}

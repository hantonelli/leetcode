package countofsmallernumbersafterself

func countSmaller2(nums []int) []int {
	viewed := map[int]int{}
	res := make([]int, len(nums))
	var n, count int
	for i := len(nums) - 1; 0 <= i; i-- {
		n = nums[i]
		count = 0
		for k, v := range viewed {
			if k < n {
				count += v
			}
		}
		res[i] = count
		viewed[n]++
	}
	return res
}

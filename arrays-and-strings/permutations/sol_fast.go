package permutations

func permuteFast(nums []int) [][]int {
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		var rest = make([]int, len(nums))
		copy(rest, nums)
		ans = append(ans, gen([]int{nums[i]}, append(rest[:i], rest[i+1:]...))...)
	}
	return ans
}

func gen(curr []int, rest []int) [][]int {
	if len(rest) == 0 {
		return [][]int{curr}
	}
	var ans [][]int
	for i := 0; i < len(rest); i++ {
		var currCopy = make([]int, len(curr))
		copy(currCopy, curr)
		currCopy = append(curr, rest[i])
		var buf = make([]int, len(rest))
		copy(buf, rest)
		ans = append(ans, gen(currCopy, append(buf[:i], buf[i+1:]...))...)

	}
	return ans
}

package sum3

// import "sort"

// func threeSumFast(nums []int) [][]int {
// 	var ans [][]int

// 	sort.Ints(nums)

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > 0 || (i > 0 && nums[i] == nums[i-1]) {
// 			continue
// 		}

// 		j := i + 1
// 		k := len(nums) - 1

// 		for j < k {
// 			sum := nums[i] + nums[j] + nums[k]
// 			if sum < 0 {
// 				j++
// 			} else if sum > 0 {
// 				k--
// 			} else {
// 				ans = append(ans, []int{nums[i], nums[j], nums[k]})

// 				for j < k && nums[j] == nums[j+1] {
// 					j++
// 				}
// 				for j < k && nums[k] == nums[k-1] {
// 					k--
// 				}

// 				j++
// 				k--
// 			}
// 		}
// 	}

// 	return ans
// }

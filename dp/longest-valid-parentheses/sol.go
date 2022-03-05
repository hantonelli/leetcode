package longestvalidparentheses

import ds "github.com/hantonelli/leetcode/data-structures"

func longestValidParentheses(s string) int {
	if len(s) < 1 {
		return 0
	}
	maxLen := 0
	stack := ds.NewStack3[int]()
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				// The last ( is pushed, as we are going to start to count from
				// there to count the continuous occurrences
				// (and then there is no need to add a +1 to the maxLen calculation)
				stack.Push(i)
			} else {
				maxLen = ds.Max(maxLen, i-stack.Peek())
			}
		}
	}
	return maxLen
}

// func longestValidParentheses(s string) int {
// 	if len(s) < 1 {
// 		return 0
// 	}
// 	maxLen := 0
// 	stack := []int{-1}
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			stack = append(stack, i)
// 		} else {
// 			if len(stack) != 0 {
// 				stack = stack[:len(stack)-1]
// 			}
// 			if len(stack) == 0 {
// 				stack = append(stack, i)
// 			} else {
// 				maxLen = max(maxLen, i-stack[len(stack)-1])
// 			}
// 		}
// 	}
// 	return maxLen
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

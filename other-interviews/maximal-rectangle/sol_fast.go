package maximalrectangle

import (
	ds "github.com/hantonelli/leetcode/data-structures"
)

// Get the maximum area in a histogram given its heights
func leetcode84(heights []int) int {
	stack := ds.NewStack()
	stack.Push(-1)
	maxarea := 0
	for i := 0; i < len(heights); i++ {
		for stack.Peek() != -1 && heights[stack.Peek().(int)] >= heights[i] {
			maxarea = max(maxarea, heights[stack.Pop().(int)]*(i-stack.Peek().(int)-1))
		}
		stack.Push(i)
	}
	for stack.Peek() != -1 {
		maxarea = max(maxarea, heights[stack.Pop().(int)]*(len(heights)-stack.Peek().(int)-1))
	}
	return maxarea
}

func maximalRectangleFast(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxarea := 0
	dp := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// update the state of this row's histogram using the last row's histogram
			// by keeping track of the number of consecutive ones
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		// update maxarea with the maximum area from this row's histogram
		maxarea = max(maxarea, leetcode84(dp))
	}
	return maxarea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

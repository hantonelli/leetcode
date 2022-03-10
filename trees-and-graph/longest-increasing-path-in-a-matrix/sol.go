package longestincreasingpathinamatrix

func longestIncreasingPath(matrix [][]int) int {
	memo := map[[2]int]int{}
	res := 0
	for col := 0; col < len(matrix); col++ {
		for row := 0; row < len(matrix[col]); row++ {
			res = max(res, dp([2]int{col, row}, matrix, memo))
		}
	}
	return res
}

func dp(pos [2]int, m [][]int, memo map[[2]int]int) int {
	if v, ok := memo[pos]; ok {
		return v
	}
	var res int
	col := pos[0]
	row := pos[1]
	if col+1 < len(m) && m[col][row] < m[col+1][row] {
		res = max(res, dp([2]int{col + 1, row}, m, memo))
	}
	if row+1 < len(m[0]) && m[col][row] < m[col][row+1] {
		res = max(res, dp([2]int{col, row + 1}, m, memo))
	}
	if 0 <= col-1 && m[col][row] < m[col-1][row] {
		res = max(res, dp([2]int{col - 1, row}, m, memo))
	}
	if 0 <= row-1 && m[col][row] < m[col][row-1] {
		res = max(res, dp([2]int{col, row - 1}, m, memo))
	}
	memo[pos] = res + 1
	return res + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

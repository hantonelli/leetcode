package longestincreasingpathinamatrix

func longestIncreasingPathFastest(matrix [][]int) int {
	m, n, res := len(matrix), len(matrix[0]), 1

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := dfsFastest(matrix, dp, i, j)
			if res < count {
				res = count
			}
		}
	}

	return res
}

func dfsFastest(matrix, dp [][]int, i, j int) int {
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	m, n, count := len(matrix), len(matrix[0]), 1

	if i-1 >= 0 && matrix[i][j] < matrix[i-1][j] {
		count = max(count, 1+dfsFastest(matrix, dp, i-1, j))
	}
	if i+1 < m && matrix[i][j] < matrix[i+1][j] {
		count = max(count, 1+dfsFastest(matrix, dp, i+1, j))
	}
	if j-1 >= 0 && matrix[i][j] < matrix[i][j-1] {
		count = max(count, 1+dfsFastest(matrix, dp, i, j-1))
	}
	if j+1 < n && matrix[i][j] < matrix[i][j+1] {
		count = max(count, 1+dfsFastest(matrix, dp, i, j+1))
	}
	dp[i][j] = count
	return count
}

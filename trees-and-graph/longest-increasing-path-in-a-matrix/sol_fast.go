package longestincreasingpathinamatrix

// From a position, we can just go to the bigger positions.
// That allows that a position can be cached, without tracking the previous
// visited positions, as we cannot visit them as they have lower values the
// current position.

var dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func longestIncreasingPathFast(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfsFast(i, j, m, n, matrix, cache))
		}
	}
	return ans
}

func dfsFast(i, j, m, n int, matrix [][]int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	for _, d := range dirs {
		x := i + d[0]
		y := j + d[1]
		if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] > matrix[i][j] {
			cache[i][j] = max(cache[i][j], dfsFast(x, y, m, n, matrix, cache))
		}
	}
	cache[i][j]++
	return cache[i][j]
}

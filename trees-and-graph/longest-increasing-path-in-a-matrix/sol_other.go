package longestincreasingpathinamatrix

func longestIncreasingPath2(matrix [][]int) int {
	var res, curr, leng, maxLeng int
	var found bool
	visited := make(map[[2]int]bool, len(matrix)*len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			curr, leng = dfs(i, j, matrix, visited)
			if !found || res < curr {
				found = true
				res = curr
				maxLeng = leng
			}
		}
	}
	return maxLeng
}

func dfs(i, j int, matrix [][]int, visited map[[2]int]bool) (int, int) {
	var res, curr, leng, maxLeng int
	var found bool
	visited[[2]int{i, j}] = true
	if i+1 < len(matrix) && !visited[[2]int{i + 1, j}] {
		visited[[2]int{i + 1, j}] = true
		curr, leng = dfs(i+1, j, matrix, visited)
		if !found || res < curr {
			found = true
			res = curr
			maxLeng = leng
		}
		visited[[2]int{i + 1, j}] = false
	}
	if j+1 < len(matrix[0]) && !visited[[2]int{i, j + 1}] {
		visited[[2]int{i, j + 1}] = true
		curr, leng = dfs(i, j+1, matrix, visited)
		if !found || res < curr {
			found = true
			res = curr
			maxLeng = leng
		}
		visited[[2]int{i, j + 1}] = false
	}
	if 0 < i-1 && !visited[[2]int{i - 1, j}] {
		visited[[2]int{i - 1, j}] = true
		curr, leng = dfs(i-1, j, matrix, visited)
		if !found || res < curr {
			found = true
			res = curr
			maxLeng = leng
		}
		visited[[2]int{i - 1, j}] = false
	}
	if 0 < j-1 && !visited[[2]int{i, j - 1}] {
		visited[[2]int{i, j - 1}] = true
		curr, leng = dfs(i, j-1, matrix, visited)
		if !found || res < curr {
			found = true
			res = curr
			maxLeng = leng
		}
		visited[[2]int{i, j - 1}] = false
	}
	visited[[2]int{i, j}] = false
	if found {
		return matrix[i][j] + res, maxLeng + 1
	}
	return matrix[i][j], 1
}

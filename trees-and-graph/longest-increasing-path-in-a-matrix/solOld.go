package longestincreasingpathinamatrix

func longestIncreasingPathOld(matrix [][]int) int {
	visited := map[[2]int]bool{}
	res := 0
	for col := 0; col < len(matrix); col++ {
		for row := 0; row < len(matrix[col]); row++ {
			res = max(res, dpOld([2]int{col, row}, matrix, visited))
		}
	}
	return res
}

func dpOld(pos [2]int, m [][]int, v map[[2]int]bool) int {
	v[pos] = true
	res := len(v)
	col := pos[0]
	row := pos[1]
	if col+1 < len(m) && !v[[2]int{col + 1, row}] && m[col][row] < m[col+1][row] {
		res = max(res, dpOld([2]int{col + 1, row}, m, v))
	}
	if row+1 < len(m[0]) && !v[[2]int{col, row + 1}] && m[col][row] < m[col][row+1] {
		res = max(res, dpOld([2]int{col, row + 1}, m, v))
	}
	if 0 <= col-1 && !v[[2]int{col - 1, row}] && m[col][row] < m[col-1][row] {
		res = max(res, dpOld([2]int{col - 1, row}, m, v))
	}
	if 0 <= row-1 && !v[[2]int{col, row - 1}] && m[col][row] < m[col][row-1] {
		res = max(res, dpOld([2]int{col, row - 1}, m, v))
	}
	delete(v, pos)
	return res
}

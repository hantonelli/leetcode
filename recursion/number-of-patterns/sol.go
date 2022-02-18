package numberofpatterns

func numberOfPatterns(m int, n int) int {
	visited := make(map[[2]int]bool, 9)
	count := 0
	for k := n; k < m; k++ {
		count += processWithLength(k, &visited)
	}
	return count
}

func processWithLength(k int, visited map[[2]int]bool) int {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			count += process(i, j, 1, k, visited)
		}
	}
	return count
}

func process(i, j, curr, k int, visited map[[2]int]bool) int {
	if curr == k {
		return 1
	}
	count := 0
	visited[[2]int{i, j}] = true
	availableOptions := getAvailable(i, j, visited)
	for _, o := range availableOptions {
		visited[[2]int{o[0], o[1]}] = true
		count += process(o[0], o[1], curr+1, k, visited)
		visited[[2]int{o[0], o[1]}] = false
	}
	visited[[2]int{i, j}] = false
	return 0
}

func getAvailable(i, j int, visited map[[2]int]bool) [][2]int {
	var res [][2]int
	if 0 <= i-1 && !visited[[2]int{i - 1, j}] {
		res = append(res, [2]int{i - 1, j})
	}
	if 0 <= j-1 && !visited[[2]int{i, j - 1}] {
		res = append(res, [2]int{i, j - 1})
	}

	if 0 <= j-2 && visited[[2]int{i, j - 1}] && !visited[[2]int{i, j - 2}] {
		res = append(res, [2]int{i, j - 2})
	}
	if 0 <= i-2 && visited[[2]int{i - 1, j}] && !visited[[2]int{i - 2, j}] {
		res = append(res, [2]int{i - 2, j})
	}

	if i+1 < 3 && !visited[[2]int{i + 1, j}] {
		res = append(res, [2]int{i + 1, j})
	}
	if j+1 < 3 && !visited[[2]int{i, j + 1}] {
		res = append(res, [2]int{i, j + 1})
	}

	if i+2 < 3 && visited[[2]int{i + 1, j}] && !visited[[2]int{i + 2, j}] {
		res = append(res, [2]int{i + 2, j})
	}
	if j+2 < 3 && visited[[2]int{i, j + 1}] && !visited[[2]int{i, j + 2}] {
		res = append(res, [2]int{i, j + 2})
	}

	if j+2 < 3 && visited[[2]int{i, j + 1}] && !visited[[2]int{i, j + 2}] {
		res = append(res, [2]int{i, j + 2})
	}

	return res
}

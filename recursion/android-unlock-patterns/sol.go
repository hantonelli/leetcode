package androidunlockpatterns

func numberOfPatterns(min int, max int) int {
	count := 0
	visited := map[int]bool{}
	for pos := 1; pos <= 9; pos++ {
		count += dp(min, max, pos, visited)
	}
	return count
}

func dp(min int, max int, pos int, visited map[int]bool) int {
	visited[pos] = true
	if len(visited) == max {
		delete(visited, pos)
		return 1
	}
	count := 0
	if min <= len(visited) {
		count++
	}
	nexts := getAvailableNext(pos, visited)
	for _, next := range nexts {
		count += dp(min, max, next, visited)
	}
	delete(visited, pos)
	return count
}

// 1 2 3
// 4 5 6
// 7 8 9

var options = map[int][]int{
	0:  {1, 2, 3, 5, 4, 6, 7, 8, 9},
	10: {2, 4, 5},
	20: {1, 3, 4, 5, 6},
	30: {2, 5, 6},
	40: {1, 2, 5, 7, 8},
	50: {1, 2, 3, 4, 6, 7, 8, 9},
	60: {2, 3, 5, 8, 9},
	70: {4, 5, 8},
	80: {4, 5, 6, 7, 9},
	90: {8, 5, 6},
	11: {2, 4, 6, 8, 9},
	31: {2, 4, 6, 7, 8},
	71: {2, 3, 4, 6, 8},
	91: {1, 2, 4, 6, 8},
}

func getAvailableNext(pos int, visited map[int]bool) []int {
	var opts []int
	if pos == 5 || (visited[5] && (pos == 2 || pos == 4 || pos == 6 || pos == 8)) {
		opts = options[0]
	} else {
		if !visited[5] {
			opts = options[pos*10]
		} else {
			opts = options[pos*10+1]
		}
		switch pos {
		case 1:
			if visited[2] && !visited[3] {
				opts = append(opts, 3)
			}
			if visited[4] && !visited[7] {
				opts = append(opts, 7)
			}
		case 3:
			if visited[2] && !visited[1] {
				opts = append(opts, 1)
			}
			if visited[6] && !visited[9] {
				opts = append(opts, 9)
			}
		case 7:
			if visited[4] && !visited[1] {
				opts = append(opts, 1)
			}
			if visited[8] && !visited[9] {
				opts = append(opts, 9)
			}
		case 9:
			if visited[6] && !visited[3] {
				opts = append(opts, 3)
			}
			if visited[8] && !visited[7] {
				opts = append(opts, 7)
			}
		}
	}

	var res []int
	for _, o := range opts {
		if !visited[o] {
			res = append(res, o)
		}
	}
	return res
}

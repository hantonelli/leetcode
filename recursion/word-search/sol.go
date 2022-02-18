package wordsearch

func exist(board [][]byte, word string) bool {
	exists := false
	visited := make(map[[2]int]bool, len(board)*len(board[0]))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				if process(i, j, board, visited, word, 1) {
					return true
				}
			}
		}
	}
	return exists
}

func process(i, j int, board [][]byte, visited map[[2]int]bool, word string, pos int) bool {
	if pos == len(word) {
		return true
	}
	visited[[2]int{i, j}] = true

	if 0 <= i-1 && !visited[[2]int{i - 1, j}] && word[pos] == board[i-1][j] {
		visited[[2]int{i - 1, j}] = true
		if process(i-1, j, board, visited, word, pos+1) {
			return true
		}
		visited[[2]int{i - 1, j}] = false
	}
	if 0 <= j-1 && !visited[[2]int{i, j - 1}] && word[pos] == board[i][j-1] {
		visited[[2]int{i, j - 1}] = true
		if process(i, j-1, board, visited, word, pos+1) {
			return true
		}
		visited[[2]int{i, j - 1}] = false
	}
	if i+1 < len(board) && !visited[[2]int{i + 1, j}] && word[pos] == board[i+1][j] {
		visited[[2]int{i + 1, j}] = true
		if process(i+1, j, board, visited, word, pos+1) {
			return true
		}
		visited[[2]int{i + 1, j}] = false
	}
	if j+1 < len(board[0]) && !visited[[2]int{i, j + 1}] && word[pos] == board[i][j+1] {
		visited[[2]int{i, j + 1}] = true
		if process(i, j+1, board, visited, word, pos+1) {
			return true
		}
		visited[[2]int{i, j + 1}] = false
	}

	visited[[2]int{i, j}] = false
	return false
}

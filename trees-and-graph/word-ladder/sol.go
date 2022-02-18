package wordladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make(map[string]bool, len(wordList))
	diffMemo := map[string]bool{}
	res := getShorterPath(beginWord, endWord, wordList, 0, -1, visited, diffMemo)
	if res == -1 {
		return 0
	}
	return res
}

func getShorterPath(beginWord string, endWord string, wordList []string, prevMinPathSteps, lastFullPath int, visited map[string]bool, diffMemo map[string]bool) int {
	if beginWord == endWord {
		return prevMinPathSteps + 1
	}
	if lastFullPath != -1 && lastFullPath < prevMinPathSteps {
		return -1
	}
	minPathSteps := -1
	visited[beginWord] = true
	for _, word := range wordList {
		if !visited[word] && diffOne(beginWord, word, diffMemo) {
			steps := getShorterPath(word, endWord, wordList, prevMinPathSteps+1, minPathSteps, visited, diffMemo)
			if steps != -1 && (minPathSteps == -1 || steps < minPathSteps) {
				minPathSteps = steps
			}
		}
	}
	visited[beginWord] = false
	return minPathSteps
}

func diffOne(a, b string, diffMemo map[string]bool) bool {
	if a < b {
		if v, ok := diffMemo[a+b]; ok {
			return v
		}
	} else {
		if v, ok := diffMemo[b+a]; ok {
			return v
		}
	}
	var diffs int
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}
	if a < b {
		diffMemo[a+b] = diffs == 1
	} else {
		diffMemo[b+a] = diffs == 1
	}
	return diffs == 1
}

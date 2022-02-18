package wordladder

func ladderLength3(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}

	wordMap := make(map[string]bool)
	for _, s := range wordList {
		wordMap[s] = true
	}

	vis := make(map[string]bool)
	vis[beginWord] = true
	dist := 1
	for !contains(endWord, vis) && len(vis) > 0 {
		nextVis := make(map[string]bool)
		for s := range vis {
			for i := 0; i < len(s); i++ {
				for c := 'a'; c <= 'z'; c++ {
					next := s[:i] + string(c) + s[i+1:]
					if !contains(next, wordMap) {
						continue
					}

					nextVis[next] = true
					delete(wordMap, next)
				}
			}
		}
		vis = nextVis
		dist++
	}

	if contains(endWord, vis) {
		return dist
	}

	return 0
}

func contains(s string, m map[string]bool) bool {
	if _, ok := m[s]; ok {
		return true
	}

	return false
}

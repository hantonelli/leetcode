package wordladder2

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	current := [][]string{{beginWord}}
	candidates := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		candidates[w] = true
	}

	res := [][]string{}
	found := false
	for deep := 1; 0 < len(current) && !found; deep++ {
		for _, w := range current {
			delete(candidates, w[len(w)-1])
		}
		next := [][]string{}
		for _, wl := range current {
			w := wl[len(wl)-1]
			for pos := 0; pos < len(w); pos++ {
				for i := 'a'; i <= 'z'; i++ {
					possibleWord := w[:pos] + string(i) + w[pos+1:]
					if candidates[possibleWord] {
						if possibleWord == endWord {
							found = true
							r := append([]string{}, wl...)
							res = append(res, append(r, possibleWord))
						}
						n := make([]string, 0, len(wl)+1)
						n = append(n, wl...)
						next = append(next, append(n, possibleWord))
					}
				}
			}
		}
		current = next
	}
	return res
}

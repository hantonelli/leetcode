package wordladder

func ladderLength5(beginWord string, endWord string, wordList []string) int {
	wl := map[string]bool{}
	for _, w := range wordList {
		wl[w] = true
	}
	if !wl[endWord] {
		return 0
	}
	wordLen := len(beginWord)
	level := 1
	start := map[string]bool{beginWord: true}
	end := map[string]bool{endWord: true}
	delete(wl, beginWord)
	delete(wl, endWord)
	for len(start) != 0 && len(end) != 0 {
		if len(start) > len(end) {
			start, end = end, start
		}

		temp := map[string]bool{}
		for wordString, _ := range start {
			word := []byte(wordString)
			for i := 0; i < wordLen; i++ {
				var c byte
				for c = 'a'; c <= 'z'; c++ {
					word[i] = c
					w := string(word)

					if wl[w] {
						temp[w] = true
						delete(wl, w)
					} else if end[w] {
						return level + 1
					}
				}
				word[i] = wordString[i]
			}
		}
		level++
		start = temp
	}
	return 0
}

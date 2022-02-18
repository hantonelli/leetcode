package wordladder

var alphabets = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	current := []string{beginWord}
	//use a map to indicate whether an element has been searched out
	candidate := make(map[string]bool, len(wordList))
	for _, v := range wordList {
		candidate[v] = true
	}
	//BFS. if any updated word matching endWord, return the depth.
	//As it's BFS, the first found is the shortest path
	for depth := 1; len(current) > 0; depth++ {
		// should not search the same deepth nodes.
		for _, v := range current {
			delete(candidate, v)
		}

		next := []string{}
		for _, curWord := range current {
			//go through one by one
			for i := 0; i < len(curWord); i++ {
				for _, w := range alphabets {
					// get the modified word
					updated := curWord[0:i] + w + curWord[i+1:]
					// if it's one of WordList, it's one of next level word.
					if _, ok := candidate[updated]; ok {
						//if it's the endword, return current depth
						if updated == endWord {
							return depth + 1
						}
						next = append(next, updated)
					}
				}
			}
		}
		current = next
	}
	return 0
}

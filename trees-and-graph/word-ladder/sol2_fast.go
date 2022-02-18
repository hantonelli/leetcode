package wordladder

var (
	alphabetsB = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
)

func ladderLength6(beginWord string, endWord string, wordList []string) int {
	current := []string{beginWord}
	candidates := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		candidates[w] = true
	}

	var possibleWordB []byte
	var possibleWord string
	for deep := 1; 0 < len(current); deep++ {
		for _, w := range current {
			delete(candidates, w)
		}
		next := []string{}
		for _, word := range current {
			for pos := 0; pos < len(word); pos++ {
				possibleWordB = []byte(word)
				for _, ch := range alphabetsB {
					possibleWordB[pos] = ch
					possibleWord = string(possibleWordB)
					if candidates[possibleWord] {
						if possibleWord == endWord {
							return deep + 1
						}
						next = append(next, possibleWord)
					}
				}
			}
		}
		current = next
	}
	return 0
}

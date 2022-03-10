package wordladder2

type item struct {
	word  string
	paths [][]string
}

func findLaddersFast(beginWord string, endWord string, wordList []string) [][]string {
	hasBeginWord := false
	for _, w := range wordList {
		if w == beginWord {
			hasBeginWord = true
			break
		}
	}
	if !hasBeginWord {
		wordList = append(wordList, beginWord)
	}

	g := getGraph(wordList)

	if _, found := g[endWord]; !found {
		return nil
	}

	queue := []item{
		{
			word:  beginWord,
			paths: [][]string{{beginWord}},
		},
	}

	visited := map[string]bool{beginWord: true}

	for len(queue) > 0 {
		next := map[string][][]string{}
		for _, q := range queue {
			paths := q.paths
			if q.word == endWord {
				return paths
			}
			for _, nt := range g[q.word] {
				if visited[nt] {
					continue
				}
				for _, path := range paths {
					tmp := make([]string, len(path))
					copy(tmp, path)
					tmp = append(tmp, nt)
					next[nt] = append(next[nt], tmp)
				}
			}
		}
		queue = []item{}
		for word, paths := range next {
			visited[word] = true
			queue = append(queue, item{
				word:  word,
				paths: paths,
			})
		}
	}
	return nil
}

func getGraph(words []string) map[string][]string {
	m := map[string][]string{}
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			w1, w2 := words[i], words[j]
			if diffByOne(w1, w2) {
				m[w1] = append(m[w1], w2)
				m[w2] = append(m[w2], w1)
			}
		}
	}
	return m
}

func diffByOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		} else {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}

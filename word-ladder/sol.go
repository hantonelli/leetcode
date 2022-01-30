package main

import "container/list"

// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	if beginWord == endWord {
// 		return 1
// 	}
// 	lastCount := 0
// 	for i := range wordList {
// 		w := wordList[i]
// 		if beginWord != w && isOneCharDiff(beginWord, w) {
// 			newWordList := make([]string, 0, len(wordList)-1)
// 			newWordList = append(newWordList, wordList[:i]...)
// 			newWordList = append(newWordList, wordList[i+1:]...)
// 			newCount := ladderLength(w, endWord, newWordList)
// 			if newCount != 0 {
// 				if lastCount == 0 || newCount < lastCount {
// 					lastCount = newCount
// 				}
// 			}
// 		}
// 	}
// 	if lastCount == 0 {
// 		return 0
// 	}
// 	return lastCount + 1
// }

type Queue *[]string

func (q Queue) Push(x string) {
	q = append(&q, x)
}
func (q Queue) Pop() {
	if len(q) > 0 {
		q[0] = ""
		q = q[1:]
	}
}
func (q Queue) Len() int {
	return len(q)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := list.New()

	for i := range wordList {
		queue.PushBack(wordList[i])
	}

	lastCount := 0
	for i := range wordList {
		w := wordList[i]
		if beginWord != w && isOneCharDiff(beginWord, w) {
			newWordList := make([]string, 0, len(wordList)-1)
			newWordList = append(newWordList, wordList[:i]...)
			newWordList = append(newWordList, wordList[i+1:]...)
			newCount := ladderLength(w, endWord, newWordList)
			if newCount != 0 {
				if lastCount == 0 || newCount < lastCount {
					lastCount = newCount
				}
			}
		}
	}
	if lastCount == 0 {
		return 0
	}
	return lastCount + 1
}

func isOneCharDiff(x, y string) bool {
	hasOneDiff := false
	for i := range x {
		if x[i] != y[i] {
			if hasOneDiff {
				return false
			}
			hasOneDiff = true
		}
	}
	return hasOneDiff
}

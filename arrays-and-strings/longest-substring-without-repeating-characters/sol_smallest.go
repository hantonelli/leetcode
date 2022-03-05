package main

func lengthOfLongestSubstringSmallest(s string) int {
	maxLen, start := 0, 0
	// upgrade1: since string only has 128 valid chars, use 128 char array instead of map[byte]int
	charToIndex := make([]int, 128)
	// since 0 is a valid index, init array (use as map) to -1
	for i := range charToIndex {
		charToIndex[i] = -1
	}
	for end := range s {
		if i := charToIndex[s[end]]; i != -1 {
			if i >= start {
				start = i + 1
			}
		}
		charToIndex[s[end]] = end
		if maxLen < (end-start)+1 {
			maxLen = end - start + 1
		}
	}
	return maxLen
}

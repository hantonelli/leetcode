package main

import (
	"github.com/hantonelli/leetcode/ds"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := 0
	lastChars := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if pos, ok := lastChars[s[i]]; ok {
			start = ds.Max(start, pos+1)
		}
		lastChars[s[i]] = i
		maxLen = ds.Max(maxLen, i-start+1)
	}
	return maxLen
}

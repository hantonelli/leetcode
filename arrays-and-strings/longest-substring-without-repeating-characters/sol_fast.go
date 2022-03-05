package main

func lengthOfLongestSubstringFast(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	lastSeenCharIdx := map[byte]int{}
	idxHead := 0
	longest := 0
	for i := 0; i < len(s); i++ {
		if idx, ok := lastSeenCharIdx[s[i]]; ok {
			idxHead = idx
			lastSeenCharIdx[s[i]] = i + 1
			continue
		}
		lastSeenCharIdx[s[i]] = i + 1
		if longest < i+1-idxHead {
			longest = i + 1 - idxHead
		}
	}
	return longest
}

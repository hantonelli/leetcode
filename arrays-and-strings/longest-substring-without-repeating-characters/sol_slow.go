package main

func lengthOfLongestSubstringSlow(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	longest := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if checkUnique(s, i, j) {
				if longest < j-i {
					longest = j - i
				}
			}
		}
	}
	return longest
}

func checkUnique(s string, start, end int) bool {
	chars := map[byte]bool{}
	for i := start; i < end; i++ {
		if _, exists := chars[s[i]]; exists {
			return false
		}
		chars[s[i]] = true
	}
	return true
}

package longest

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	maxLen := 0

	d1Count := 1
	d1Continuous := 1
	d1Char := s[0]

	d2Count := 0
	var d2Char byte
	d2Continuous := 0

	for i := 1; i < len(s); i++ {
		ch := string(s[i])
		fmt.Printf("%s", ch)
		if s[i] == d1Char {
			d1Count++
			d1Continuous++
			d2Continuous = 0
			continue
		}
		if s[i] == d2Char {
			d2Count++
			d2Continuous++
			d1Continuous = 0
			continue
		}

		if d2Count == 0 {
			d2Char = s[i]
			d2Count = 1
			d2Continuous = 1
			d1Continuous = 0
			continue
		}

		maxLen = max(maxLen, d2Count+d1Count)
		if s[i-1] == d1Char {
			d2Char = d1Char
			d2Count = d1Continuous
			d2Continuous = d1Continuous
		} else if s[i-1] == d2Char {
			// prevChar = prevChar
			d2Count = 1
			d2Continuous = 1
		}
		d1Char = s[i]
		d1Count = 1
		d1Continuous = 1
	}
	maxLen = max(maxLen, d2Count+d1Count)
	return maxLen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

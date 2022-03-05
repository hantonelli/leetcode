package longest

func lengthOfLongestSubstringTwoDistinctEasy(s string) int {
	if len(s) == 1 || len(s) == 0 {
		return len(s)
	}
	maxLen, length, countLast := 0, 0, 0
	var l1, last string = "", ""
	for index := range s {
		c := string(s[index])
		if c == last || c == l1 {
			length++
		} else {
			length = countLast + 1
		}
		if last == c {
			countLast++
		} else {
			countLast = 1
		}
		if last != c {
			l1 = last
			last = c
		}
		if maxLen < length {
			maxLen = length
		}
	}
	return maxLen
}

package longest

func lengthOfLongestSubstringTwoDistinctFast(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	maxLen, length := 0, 0
	l1, last, countLast := -1, -1, 0
	for index := range s {
		if int(s[index]) == last || int(s[index]) == l1 {
			length++
		} else {
			length = countLast + 1
		}
		if last == int(s[index]) {
			countLast++
		} else {
			countLast = 1
		}
		if last != int(s[index]) {
			l1 = last
			last = int(s[index])
		}
		if maxLen < length {
			maxLen = length
		}
	}
	return maxLen
}

package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for {
		for left < len(s) && !isLetter(s[left]) {
			left++
		}
		for 0 <= right && !isLetter(s[right]) {
			right--
		}
		if right < left {
			break
		}
		if s[right] != s[left] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

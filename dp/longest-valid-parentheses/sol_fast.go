package longestvalidparentheses

func longestValidParenthesesFast(s string) int {
	maxCount := 0
	left := 0
	right := 0

	// scan from left to right
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			left++
		} else {
			right++
		}
		if left == right && left > maxCount {
			maxCount = left
		}
		if left < right {
			left = 0
			right = 0
		}
	}

	// scan from right to left
	left = 0
	right = 0
	for i := len(s) - 1; i > 0; i-- {
		c := s[i]
		if c == '(' {
			left++
		} else {
			right++
		}
		if left == right && left > maxCount {
			maxCount = left
		}
		if left > right {
			left = 0
			right = 0
		}
	}

	return maxCount * 2
}

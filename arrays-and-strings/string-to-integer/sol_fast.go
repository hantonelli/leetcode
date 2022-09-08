package string_to_integer

import (
	"math"
	"unicode"
)

func myAtoiFast(s string) int {
	result := 0
	sign := 1
	started := false

	for i := 0; i < len(s); i++ {
		if !started {
			if rune(s[i]) == '-' {
				sign = -1
				started = true
			} else if rune(s[i]) == '+' {
				started = true
			} else if rune(s[i]) == ' ' {
				continue
			} else if unicode.IsNumber(rune(s[i])) {
				started = true
				i--
			} else {
				break
			}
		} else {
			if unicode.IsNumber(rune(s[i])) {
				result = (result * 10) + int(rune(s[i])-'0')
				if (result * sign) > math.MaxInt32 {
					return math.MaxInt32
				} else if (result * sign) < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
	}

	return result * sign
}

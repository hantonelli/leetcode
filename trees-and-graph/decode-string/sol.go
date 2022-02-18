package decodeString

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var acc strings.Builder
	var accDigit string
	var ch rune
	for i := 0; i < len(s); i++ {
		ch = rune(s[i])
		if isLetter(ch) {
			acc.WriteRune(ch)
		}
		if isDigit(ch) {
			accDigit += string(ch)
		}
		if ch == '[' {
			open := 1
			i++
			iStart := i
			for {
				if s[i] == '[' {
					open++
				}
				if s[i] == ']' {
					open--
				}
				if open == 0 {
					break
				}
				i++
			}
			toProcess := string(s[iStart:i])
			inner := decodeString(toProcess)
			k, _ := strconv.Atoi(accDigit)
			acc.WriteString(strings.Repeat(inner, k))
			accDigit = ""
		}
	}
	return acc.String()
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z'
}

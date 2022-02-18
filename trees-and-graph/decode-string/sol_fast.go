package decodeString

import (
	"strconv"
	"strings"
)

func decodeString3(s string) string {
	var acc strings.Builder
	var accDigit strings.Builder
	var open, iStart, k, i int
	var ch rune
	for i = 0; i < len(s); i++ {
		ch = rune(s[i])
		if 'a' <= ch && ch <= 'z' {
			acc.WriteRune(ch)
		}
		if '0' <= ch && ch <= '9' {
			accDigit.WriteRune(ch)
		}
		if ch == '[' {
			open = 1
			i++
			iStart = i
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
			k, _ = strconv.Atoi(accDigit.String())
			acc.WriteString(strings.Repeat(decodeString(string(s[iStart:i])), k))
			accDigit.Reset()
		}
	}
	return acc.String()
}

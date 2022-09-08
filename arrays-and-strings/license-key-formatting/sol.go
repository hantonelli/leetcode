package license_key_formatting

import (
	"math"
	"unicode"
)

func licenseKeyFormatting(s string, k int) string {
	res := make([]rune, len(s)+int(math.Ceil(float64(len(s))/float64(k))))
	resIdx := len(res) - 1
	chLen := k
	for i := len(s) - 1; 0 <= i; i-- {
		if s[i] == '-' {
			continue
		}
		res[resIdx] = unicode.ToUpper(rune(s[i]))
		resIdx--
		chLen--
		if chLen == 0 && 0 < resIdx {
			res[resIdx] = '-'
			resIdx--
			chLen = k
		}
	}
	if resIdx+1 == len(res) {
		return ""
	}
	if res[resIdx+1] == '-' {
		return string(res[resIdx+2:])
	}
	return string(res[resIdx+1:])
}

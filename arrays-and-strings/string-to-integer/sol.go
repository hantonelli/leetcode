package string_to_integer

import (
	"math"
	"math/big"
)

var (
	min = big.NewInt(int64(-math.Pow(2, 31)))
	max = big.NewInt(int64(math.Pow(2, 31) - 1))
)

func myAtoi(s string) int {
	signViewed, digitStart := false, false
	start, end := 0, len(s)
	for i, ch := range []rune(s) {
		if ch == ' ' && !digitStart {
			start++
			continue
		}
		if !signViewed && !digitStart && (ch == '-' || ch == '+') {
			signViewed = true
			continue
		}
		if '0' <= ch && ch <= '9' {
			if !digitStart {
				digitStart = true
			}
		} else {
			end = i
			break
		}
	}
	res, ok := new(big.Int).SetString(s[start:end], 10)
	if !ok {
		return 0
	}
	if 0 < min.Cmp(res) {
		res = min
	}
	if max.Cmp(res) < 0 {
		res = max
	}
	return int(res.Int64())
}

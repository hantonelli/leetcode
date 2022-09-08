package add_binary

import (
	"github.com/hantonelli/leetcode/ds"
)

func addBinary(a string, b string) string {
	resLen := ds.Max(len(a), len(b)) + 1
	res := make([]rune, resLen)
	var rest int
	iA, iB, iJ := len(a)-1, len(b)-1, resLen-1
	finished := false
	for !finished {
		sum := 0
		if iA < 0 && iB < 0 {
			finished = true
			sum = rest
		} else if 0 <= iA && 0 <= iB {
			sum = int(a[iA]-'0') + int(b[iB]-'0') + rest
			iA--
			iB--
		} else if 0 <= iA {
			sum = int(a[iA]-'0') + rest
			iA--
		} else {
			sum = int(b[iB]-'0') + rest
			iB--
		}
		switch sum {
		case 0:
			res[iJ] = '0'
			rest = 0
		case 1:
			res[iJ] = '1'
			rest = 0
		case 2:
			res[iJ] = '0'
			rest = 1
		case 3:
			res[iJ] = '1'
			rest = 1
		}
		iJ--
	}

	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}

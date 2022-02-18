package decodeString

import "strconv"

func decodeString2(s string) string {
	var repStack []int
	var charStack [][]byte

	var res []byte
	num := ""
	for i := 0; i < len(s); i++ {
		// fmt.Println(repStack, charStack, res, num)
		if s[i] >= '0' && s[i] <= '9' {
			num += string(s[i])
			continue
		}

		if s[i] == '[' {
			v, _ := strconv.Atoi(num)
			num = ""
			repStack = append(repStack, v)
			charStack = append(charStack, []byte{})
			continue
		}

		if s[i] == ']' {
			chars := charStack[len(charStack)-1]
			charStack = charStack[:len(charStack)-1]
			rep := repStack[len(repStack)-1]
			repStack = repStack[:len(repStack)-1]

			var iRes []byte
			for i := 0; i < rep; i++ {
				for j := 0; j < len(chars); j++ {
					iRes = append(iRes, chars[j])
				}
			}

			if len(charStack) > 0 {
				charStack[len(charStack)-1] = append(charStack[len(charStack)-1], iRes...)
			} else {
				res = append(res, iRes...)
			}

			continue
		}

		if len(charStack) > 0 {
			chars := charStack[len(charStack)-1]
			chars = append(chars, s[i])
			charStack[len(charStack)-1] = chars
			continue
		}

		res = append(res, s[i])
	}

	return string(res)
}

package main

import "fmt"

func expressiveWords2(s string, words []string) int {
	distinct := make([]byte, 0)
	counts := make([]int, 0)

	cur := s[0]
	count := 1

	for i := 1; i < len(s); i++ {
		if cur != s[i] {
			distinct = append(distinct, byte(cur))
			counts = append(counts, count)
			count = 1
			cur = s[i]
		} else {
			count++
		}
	}

	distinct = append(distinct, byte(cur))
	counts = append(counts, count)

	res := 0

	for _, w := range words {
		if len(w) < len(distinct) || len(w) > len(s) {
			continue
		}

		wc := w[0]
		count := 1
		p := 0
		ok := true
		for i := 1; i < len(w); i++ {
			if wc != w[i] {
				if byte(wc) != distinct[p] {
					ok = false
					break
				}
				if counts[p] <= 2 && count != counts[p] {
					ok = false
					break
				}

				if counts[p] < count {
					ok = false
					break
				}
				p++
				if p == len(distinct) {
					ok = false
					break
				}
				count = 1
				wc = w[i]
			} else {
				count++
			}
		}
		if !ok {
			continue
		}

		if byte(wc) != distinct[p] {
			ok = false
		}
		if counts[p] <= 2 && count != counts[p] {
			ok = false
		}

		if counts[p] < count {
			ok = false
		}

		if p != len(distinct)-1 {
			ok = false
		}

		if ok {
			fmt.Println(w)
			res++
		}
	}

	return res
}

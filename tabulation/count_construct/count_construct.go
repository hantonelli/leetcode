package main

import "strings"

func CountConstruct(target string, wordBank []string) int {
	if target == "" {
		return 1
	}

	values := make([]int, len(target)+1)
	values[0] = 1

	for i := 0; i < len(target)+1; i++ {
		if 0 < values[i] {
			for _, word := range wordBank {
				if strings.Index(target[i:], word) == 0 {
					if i+len(word) < len(target)+1 {
						values[i+len(word)] += values[i]
					}
				}
			}
		}
	}

	return values[len(target)]
}

package main

import "strings"

func CanConstruct(target string, wordBank []string) bool {
	if target == "" {
		return true
	}

	values := make([]bool, len(target)+1)
	values[0] = true

	for i := 0; i < len(target)+1; i++ {
		if values[i] {
			for _, word := range wordBank {
				if strings.Index(target[i:], word) == 0 {
					if i+len(word) < len(target)+1 {
						values[i+len(word)] = true
					}
				}
			}
		}
	}

	return values[len(target)]
}

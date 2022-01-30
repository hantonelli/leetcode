package main

import "strings"

func CountConstruct(target string, wordBank []string) int {
	memo := map[string]int{}
	return countConstruct(target, wordBank, memo)
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	if v, ok := memo[target]; ok {
		return v
	}
	if target == "" {
		return 1
	}

	count := 0
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			count += countConstruct(target[len(word):], wordBank, memo)
		}
	}

	memo[target] = count
	return count
}

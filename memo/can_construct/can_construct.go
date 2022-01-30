package main

import "strings"

func CanConstruct(target string, wordBank []string) bool {
	memo := map[string]bool{}
	return canConstruct(target, wordBank, memo)
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if v, ok := memo[target]; ok {
		return v
	}
	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			if canConstruct(target[len(word):], wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

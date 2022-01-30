package main

import "strings"

func AllConstruct(target string, wordBank []string) [][]string {
	memo := map[string][][]string{}
	return allConstruct(target, wordBank, memo)
}

func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if v, ok := memo[target]; ok {
		return v
	}
	if target == "" {
		return [][]string{{}}
	}

	convinations := [][]string{}
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			sufix := target[len(word):]
			sufixWays := allConstruct(sufix, wordBank, memo)
			for _, arr := range sufixWays {
				convinations = append(convinations, append(arr, word))
			}
		}
	}

	memo[target] = convinations
	return convinations
}

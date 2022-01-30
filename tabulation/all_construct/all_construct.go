package main

import "strings"

func AllConstruct(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	values := make([][][]string, len(target)+1)
	values[0] = [][]string{{}}

	for i := 0; i < len(target)+1; i++ {
		if values[i] != nil {
			for _, word := range wordBank {
				newTarget := target[i:]
				if strings.Index(newTarget, word) == 0 {
					if i+len(word) < len(target)+1 {
						newVal := values[i+len(word)]
						if newVal == nil {
							newVal = [][]string{}
						}
						for _, arr := range values[i] {
							newVal = append(newVal, append(arr, word))
						}
						values[i+len(word)] = newVal
					}
				}
			}
		}
	}

	return values[len(target)]
}

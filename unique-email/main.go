package main

import (
	"fmt"
	"strings"
)

// 23:30
// 24:19

func numUniqueEmails(emails []string) int {
	var b strings.Builder
	b.Grow(256)
	hasPlusAppear := false

	var domain string
	var i, j, domainID int
	var r rune
	var domainExists, emailExists bool
	var domainsMap map[string]int = make(map[string]int)
	var emailsMap map[string]bool = make(map[string]bool)

	for i = range emails {
		for j, r = range emails[i] {
			if r == '.' {
				continue
			}
			if r == '+' {
				hasPlusAppear = true
				continue
			}
			if !hasPlusAppear && emails[i][j] != '@' {
				b.WriteRune(r)
			}

			if emails[i][j] == '@' {
				domain = emails[i][(j + 1):]
				domainID, domainExists = domainsMap[domain]
				if !domainExists {
					domainID = len(domainsMap) + 1
					domainsMap[domain] = domainID
				}
				if _, emailExists = emailsMap[fmt.Sprintf("%v%v", b.String(), domainID)]; !emailExists {
					emailsMap[fmt.Sprintf("%v%v", b.String(), domainID)] = true
				}
				break
			}
		}
		b.Reset()
		hasPlusAppear = false
	}
	return len(emailsMap)
}

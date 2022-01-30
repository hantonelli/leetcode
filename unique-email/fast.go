package main

// 24:19
// 24:

func numUniqueEmails2(emails []string) int {
	emailsMap := map[string]struct{}{}
	buf := make([]byte, 100)

	for _, e := range emails {
		emailsMap[clearEmail(e, buf)] = struct{}{}
	}

	return len(emailsMap)
}

func clearEmail(email string, buf []byte) string {
	i, idx := 0, 0
	hasPlus := false

	for ; i < len(email); i++ {
		if email[i] == '.' {
			continue
		}
		if email[i] == '+' {
			hasPlus = true
		}
		if !hasPlus && email[i] != '@' {
			buf[idx] = email[i]
			idx++
		}
		if email[i] == '@' {
			break
		}
	}
	for ; i < len(email); i++ {
		buf[idx] = email[i]
		idx++
	}
	return string(buf[:idx])
}

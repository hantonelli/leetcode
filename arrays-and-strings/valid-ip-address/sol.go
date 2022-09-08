package validate_ip_address

import (
	"strconv"
	"strings"
)

const (
	NEITHER = "Neither"
	IPV4    = "IPv4"
	IPV6    = "IPv6"
)

func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		parts := strings.Split(queryIP, ".")
		if len(parts) != 4 {
			return NEITHER
		}
		for _, p := range parts {
			if !isValidV4Part(p) {
				return NEITHER
			}
		}
		return IPV4
	} else if strings.Contains(queryIP, ":") {
		parts := strings.Split(queryIP, ":")
		if len(parts) != 8 {
			return NEITHER
		}
		for _, p := range parts {
			if !isValidV6Part(p) {
				return NEITHER
			}
		}
		return IPV6
	}
	return NEITHER
}

func isValidV4Part(part string) bool {
	if len(part) == 0 || 4 <= len(part) {
		return false
	}
	if part[0] == '0' && len(part) != 1 {
		return false
	}
	partInt, err := strconv.ParseInt(part, 10, 32)
	if err != nil {
		return false
	}
	return 0 <= partInt && partInt <= 255
}

func isValidV6Part(part string) bool {
	if 0 == len(part) || 5 <= len(part) {
		return false
	}
	for _, ch := range part {
		validChar := '0' <= ch && ch <= '9' || 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F'
		if !validChar {
			return false
		}
	}
	return true
}

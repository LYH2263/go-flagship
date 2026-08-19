package flag

import (
	"strings"
	"unicode"
)

func NormalizeKey(key string) string {
	return strings.TrimSpace(strings.ToLower(key))
}

func IsValidKey(key string) bool {
	if key == "" || len(key) > 128 {
		return false
	}
	for _, r := range key {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' || r == '.' {
			continue
		}
		return false
	}
	return true
}

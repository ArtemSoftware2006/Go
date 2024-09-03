package main

import (
	"unicode"
)

// DeleteVowels(s string) string,
func DeleteVowels(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		switch unicode.ToLower(rune(s[i])) {
		case 'a':
			continue
		case 'e':
			continue
		case 'i':
			continue
		case 'o':
			continue
		case 'u':
			continue
		}
		result += string(s[i])
	}
	return result
}

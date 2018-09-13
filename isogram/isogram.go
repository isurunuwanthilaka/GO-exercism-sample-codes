package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	m := make(map[string]int)
	for _, char := range word {
		if unicode.IsLetter(char) {
			if _, ok := m[string(char)]; ok {
				return false
			} else {
				m[string(char)] = 1
			}
		}
	}
	return true
}

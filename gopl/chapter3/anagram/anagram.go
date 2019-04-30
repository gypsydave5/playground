package main

import (
	"strings"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, r := range s1 {
		i := strings.IndexRune(s2, r)
		l := runeLength(r)
		if i == -1 {
			return false
		}
		s2 = s2[:i] + s2[i+l:]
	}

	return true
}

// runeLength returns the length of a rune in bytes
func runeLength(r rune) int {
	if r > 65535 {
		return 4
	}
	if r > 2047 {
		return 3
	}
	if r > 127 {
		return 2
	}
	return 1
}

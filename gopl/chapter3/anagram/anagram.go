package main

import (
	"strings"
	"unicode/utf8"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	r, size := utf8.DecodeRuneInString(s1)
	i := strings.IndexRune(s2, r)
	if i == -1 {
		return false
	}
	return anagram(s1[size:], s2[:i]+s2[i+size:])
}

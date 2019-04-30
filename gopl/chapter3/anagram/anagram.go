package main

import "strings"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	for _, r := range s1 {
		i := strings.IndexRune(s2, r)
		if i == -1 {
			return false
		}
		return anagram(s1[1:], s2[:i]+s2[i+1:])
	}

	return true
}

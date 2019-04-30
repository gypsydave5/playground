package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	cases := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"hello", "hello", true},
		{"hello", "", false},
		{"abc", "123", false},
		{"abc", "cba", true},
		{"star", "rats", true},
		{"🌤⛅️🌥", "🌥⛅️🌤", true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v->%#v", c.word1, c.word2), func(t *testing.T) {
			got := anagram(c.word1, c.word2)
			if got != c.want {
				t.Errorf("Expected result to be %v but got %v", c.want, got)
			}
		})
	}
}

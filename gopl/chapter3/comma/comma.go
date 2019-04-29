package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123456.123456"))
}

// Exercise 3.10
func comma(s string) string {
	var result bytes.Buffer

	if s[0] == '+' || s[0] == '-' {
		result.WriteRune(rune(s[0]))
		s = s[1:]
	}

	split := strings.SplitN(s, ".", 2)
	s = split[0]

	for i, r := range s {
		result.WriteRune(r)
		if (len(s)-i-1)%3 == 0 && i != len(s)-1 {
			result.WriteRune(',')
		}
	}
	if len(split) == 2 {
		result.WriteRune('.')
		result.WriteString(split[1])
	}
	return result.String()
}

package main

import "bytes"

// Exercise 3.10
func comma(s string) string {
	var result bytes.Buffer
	for i, r := range s {
		if i+(len(s))%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(r)
	}
	return result.String()
}

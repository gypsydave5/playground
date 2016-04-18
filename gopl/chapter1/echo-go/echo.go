package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(echo1(os.Args))
	fmt.Println(echo2(os.Args))
	fmt.Println(echo3(os.Args))
	fmt.Println(echo1_1(os.Args))
	fmt.Println(echo1_2(os.Args))
	fmt.Println(echo1_3(os.Args))
}

func echo1(a []string) string {
	var s, sep string
	for i := 1; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	return s
}

func echo2(a []string) string {
	s, sep := "", ""
	for _, arg := range a[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(a []string) string {
	return strings.Join(a[1:], " ")
}

//Exercise answers
func echo1_1(a []string) string {
	return strings.Join(a, " ")
}

func echo1_2(a []string) string {
	var s, sep string
	for i, arg := range a[1:] {
		s += sep + fmt.Sprintf("%d", i+1)
		sep = " "
		s += sep + arg
	}
	return s
}

func echo1_3(a []string) string {
	return echo1_3Join(a, " ")
}

// this is a copy of the string.Join function, with comments...
func echo1_3Join(a []string, sep string) string {
	// nothing will come of nothing...
	if len(a) == 0 {
		return ""
	}

	// only one string in the slice
	if len(a) == 1 {
		return a[0]
	}

	// total length of separators in concatenated string
	n := len(sep) * (len(a) - 1)
	// add in the length of each of the strings in a
	// this gives us the length of the output string
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	// make a slice of bytes, total length of the output string
	// (strings are just slices of bytes, donchaknow?)
	b := make([]byte, n)

	// copy the first string into b
	// copy returns the number of elements (in this case bytes) copied as an int
	bp := copy(b, a[0])

	// for all the other strings `a[1:]`
	for _, s := range a[1:] {
		// copy sep into b, starting just after byte number bp
		// and incrment the bp by the number of bytes in sep
		bp += copy(b[bp:], sep)
		// and do the same for the next string
		bp += copy(b[bp:], s)
	}

	// return the b, cast to a string
	return string(b)
}

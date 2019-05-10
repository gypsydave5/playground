// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Exercise 4.8: Modify charcount to count letters, digits, and so on in their
// Unicode categories, using functions like unicode.IsLetter.
type category struct {
	name       string
	rangeTable *unicode.RangeTable
}

func (c category) String() string {
	return c.name
}

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	categories := map[category]int{
		{"Space", unicode.Space}:       0,
		{"Punctuation", unicode.Punct}: 0,
		{"Letter", unicode.Letter}:     0,
		{"Digit", unicode.Digit}:       0,
		{"UpperCase", unicode.Upper}:   0,
		{"LowerCase", unicode.Lower}:   0,
		{"TitleCase", unicode.Title}:   0,
	}

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for c := range categories {
			if unicode.Is(c.rangeTable, r) {
				categories[c] = categories[c] + 1
			}
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncategory\tcount\n")
	for c, n := range categories {
		fmt.Printf("%v\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

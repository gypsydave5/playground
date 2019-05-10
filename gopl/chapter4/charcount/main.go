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

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like unicode.IsLetter.
	categoryNames := map[*unicode.RangeTable]string{
		unicode.Space:  "Space",
		unicode.Punct:  "Punctuation",
		unicode.Letter: "Letter",
		unicode.Digit:  "Digit",
	}

	categories := map[*unicode.RangeTable]int{
		unicode.Space:  0,
		unicode.Punct:  0,
		unicode.Letter: 0,
		unicode.Digit:  0,
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
		for rt := range categories {
			if unicode.Is(rt, r) {
				categories[rt] = categories[rt] + 1
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
	for rt, n := range categories {
		fmt.Printf("%v\t%d\n", categoryNames[rt], n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

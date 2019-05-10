// Exercise 4.9: Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before the first
// call to Scan to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	frequency := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		frequency[word] = frequency[word] + 1
	}

	fmt.Print("\nword\tcount\n")
	for word, count := range frequency {
		fmt.Printf("%s\t%d\n", word, count)
	}
}

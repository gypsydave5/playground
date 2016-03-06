// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
//
// It also now records which files were written to
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	linesInFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", linesInFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, linesInFiles)
			f.Close()
		}
	}
	printLinesRepeating(2, counts, linesInFiles)
}

func printLinesRepeating(
	threshold int,
	counts map[string]int,
	linesInFiles map[string][]string,
) {
	for line, n := range counts {
		if n >= threshold {
			fmt.Printf("%q\t%d\t%s\n", strings.Join(linesInFiles[line], ", "), n, line)
		}
	}
}

func countLines(
	f *os.File,
	counts map[string]int,
	arg string,
	linesInFiles map[string][]string,
) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !contains(linesInFiles[input.Text()], arg) {
			linesInFiles[input.Text()] = append(linesInFiles[input.Text()], arg)
		}
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
//
// It also now records which files were written to
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for _, dup := range linesRepeating(2, counts) {
		printDups(dup)
	}
}

func linesRepeating(threshold int, counts map[string]map[string]int) dupes {
	var dups dupes
	for line, fileMap := range counts {
		var files []string
		var total int
		for file, count := range fileMap {
			files = append(files, file)
			total = total + count
		}
		if total >= threshold {
			dups = append(dups, dup{
				files: files,
				line:  line,
				total: total,
			})
		}
	}
	sort.Sort(dups)
	return dups
}

func printDups(d dup) {
	fmt.Printf("%s\t%d\t%s\n", strings.Join(d.files, ", "), d.total, d.line)
}

func countLines(
	f *os.File,
	counts map[string]map[string]int,
	arg string,
) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][arg]++
	}
}
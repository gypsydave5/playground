package main

import "io"
import "bufio"

type report struct {
	count int
	files []string
}

func countRepeatLines(b io.Reader) map[string]int {
	r := make(map[string]int)
	lines := bufio.NewScanner(b)
	for lines.Scan() {
		r[lines.Text()]++
	}
	return r
}

func collateLines(c map[string]map[string]int) map[string]report {
	reports := make(map[string]report)
	for filename, repeats := range c {
		for line, count := range repeats {
			r, reportExists := reports[line]
			if reportExists {
				r.count += count
				r.files = append(r.files, filename)
				reports[line] = r
			} else {
				reports[line] = report{count, []string{filename}}
			}
		}
	}
	return reports
}

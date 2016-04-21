package main

import "io"
import "bufio"

type lineReport struct {
	line  string
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

func getReportValues(rm map[string]lineReport) []lineReport {
	var ra []lineReport
	for _, r := range rm {
		ra = append(ra, r)
	}
	return ra
}

func collateLines(c map[string]map[string]int) []lineReport {
	reports := make(map[string]lineReport)
	for filename, repeats := range c {
		for line, count := range repeats {
			r, reportExists := reports[line]
			if reportExists {
				r.line = line
				r.count += count
				r.files = append(r.files, filename)
				reports[line] = r
			} else {
				reports[line] = lineReport{line, count, []string{filename}}
			}
		}
	}
	return getReportValues(reports)
}

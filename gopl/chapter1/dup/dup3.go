package main

import (
	"bufio"
	"io"
	"sort"
)

type lineReport struct {
	line  string
	count int
	files []string
}

type lineReports []lineReport

func (rs lineReports) Len() int {
	return len(rs)
}
func (rs lineReports) Less(i, j int) bool {
	if rs[i].count > rs[j].count {
		return true
	} else if rs[i].count < rs[j].count {
		return false
	} else {
		return rs[i].line > rs[j].line
	}
}

func (rs lineReports) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func countRepeatLines(b io.Reader) map[string]int {
	r := make(map[string]int)
	lines := bufio.NewScanner(b)
	for lines.Scan() {
		r[lines.Text()]++
	}
	return r
}

func getReportValues(rm map[string]lineReport) lineReports {
	var ra []lineReport
	for _, r := range rm {
		ra = append(ra, r)
	}
	return ra
}

func updateReport(r lineReport, c int, f string) lineReport {
	var newReport lineReport
	newReport.count = c + r.count
	newReport.files = append(r.files, f)
	return newReport
}

func collateLines(c map[string]map[string]int) lineReports {
	reportMap := make(map[string]lineReport)
	for filename, repeats := range c {
		for line, count := range repeats {
			r, reportExists := reportMap[line]
			if reportExists {
				updateReport(r, count, filename)
			} else {
				reportMap[line] = lineReport{line, count, []string{filename}}
			}
		}
	}
	reports := getReportValues(reportMap)
	sort.Sort(reports)
	return reports
}

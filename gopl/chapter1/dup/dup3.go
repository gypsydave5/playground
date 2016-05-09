package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	buffers := make(map[string]io.Reader)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		buffers["stdin"] = os.Stdin
	} else {
		for _, fileName := range fileNames {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v", err)
				continue
			}
			buffers[fileName] = f
			defer f.Close()
		}
	}

	reports := processBuffers(buffers)
	fReports := filterReports(reports, func(r lineReport) bool { return r.count > 1 })
	nReports := mapReports(fReports, normalizeReportLines)
	output := formatReports(nReports, makeFormatterWithLinePad(maxLineLength(fReports)))
	fmt.Printf("%v", strings.Join(output, ""))
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
	var rs []lineReport
	for _, r := range rm {
		rs = append(rs, r)
	}
	return rs
}

func updateReport(r lineReport, c int, f string) lineReport {
	var newReport lineReport
	newReport.line = r.line
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
				reportMap[line] = updateReport(r, count, filename)
			} else {
				reportMap[line] = lineReport{line, count, []string{filename}}
			}
		}
	}
	reports := getReportValues(reportMap)
	sort.Sort(reports)
	return reports
}

func processBuffers(bm map[string]io.Reader) lineReports {
	repeatLines := make(map[string]map[string]int)
	for line, buffer := range bm {
		repeatLines[line] = countRepeatLines(buffer)
	}
	return collateLines(repeatLines)
}

func formatReports(rs lineReports, f func(lineReport) string) []string {
	result := make([]string, len(rs))
	for i, r := range rs {
		res := f(r)
		result[i] = res
	}
	return result
}

func filterReports(rs lineReports, f func(lineReport) bool) lineReports {
	result := make(lineReports, 0)
	for _, r := range rs {
		if f(r) {
			result = append(result, r)
		}
	}
	return result
}

func maxLineLength(rs lineReports) (max int) {
	for _, r := range rs {
		rLen := len([]rune(r.line))
		if rLen > max {
			max = rLen
		}
	}
	return
}

func makeFormatterWithLinePad(padTo int) func(lineReport) string {
	return func(lr lineReport) string {
		files := strings.Join(lr.files, ", ")
		return fmt.Sprintf("%-*s\t%d\t%v\n", padTo, lr.line, lr.count, files)
	}
}

func normalizeReportLines(lr lineReport) lineReport {
	normLine := strings.Replace(lr.line, "\t", "    ", -1)
	return lineReport{normLine, lr.count, lr.files}
}

func mapReports(lrs lineReports, f func(lineReport) lineReport) lineReports {
	result := make(lineReports, len(lrs))
	for i, r := range lrs {
		res := f(r)
		result[i] = res
	}
	return result
}

package main

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestMapReportsToString(t *testing.T) {
	dumbReporter := func(r lineReport) string {
		return fmt.Sprintf("%v", r.count)
	}

	rs := make(lineReports, 2)
	rs[0] = lineReport{"line1", 3, []string{"file1", "file2"}}
	rs[1] = lineReport{"line2", 66, []string{"file1"}}

	lotsOfReports := mapReport(rs, dumbReporter)

	if lotsOfReports[0] != "3" {
		t.Errorf("Expected \"3\", yet received: %v", lotsOfReports[0])
	}

	if lotsOfReports[1] != "66" {
		t.Errorf("Expected \"66\", yet received: %v", lotsOfReports[0])
	}
}

func TestFormatReport(t *testing.T) {
	lr := lineReport{"line1", 3, []string{"file1", "file2"}}
	rs := formatReport(lr)

	if rs != "'line1'\t3\tfile1, file2\n" {
		t.Errorf("Unexpected report format: %s", rs)
	}
}

func TestCountRepeatLines(t *testing.T) {
	b := strings.NewReader(`Hello
Hello
Hello
Goodbye`)
	r := countRepeatLines(b)
	if r["Hello"] != 3 {
		t.Errorf("Expected 3, but got %v", r["Hello"])
	}
}

func TestBufferMapToReports(t *testing.T) {
	bm := make(map[string]io.Reader)
	bm["fileOne"] = strings.NewReader("Hello\nHello\nHello\nGoodbye")
	bm["fileTwo"] = strings.NewReader("Hello\nGoodbye\nCiao")

	rs := processBuffers(bm)
	if rs[0].line != "Hello" {
		t.Errorf("Expected \"Hello\", got %v", rs[0].line)
	}
	if rs[2].line != "Ciao" {
		t.Errorf("Expected \"Ciao\", got %v", rs[2].line)
	}
	if rs[1].count != 2 {
		t.Errorf("Expected 2 got %v", rs[1].count)
	}
}

func TestUpdatingReport(t *testing.T) {
	count := 2
	filename := "file1"
	r1 := lineReport{"line", 3, []string{"file2"}}
	r2 := updateReport(r1, count, filename)

	if r2.count != 5 {
		t.Errorf("Expected 5, got %v", r2.count)
	}

	if r2.line != "line" {
		t.Errorf("Expected \"line\", got %v", r2.line)
	}

	eFiles := []string{"file2", "file1"}
	if !reflect.DeepEqual(r2.files, eFiles) {
		t.Errorf("Expected %v, got %v", eFiles, r2.files)
	}
}

func TestCollateLines(t *testing.T) {
	s := make(map[string]map[string]int)
	s["fileOne"] = make(map[string]int)
	s["fileTwo"] = make(map[string]int)
	s["fileOne"]["Hello"] = 3
	s["fileOne"]["Goodbye"] = 1
	s["fileTwo"]["Hello"] = 1
	s["fileTwo"]["Goodbye"] = 1
	s["fileTwo"]["Ciao"] = 1

	r := collateLines(s)
	if r[0].line != "Hello" {
		t.Errorf("Expected \"Hello\", got %v", r[0].line)
	}
	if r[0].count != 4 {
		t.Errorf("Expected 4, got %v", r[0].count)
	}
	if r[1].line != "Goodbye" {
		t.Errorf("Expected \"Goodbye\", got %v", r[1].line)
	}
}

func TestGetValuesFromReportMap(t *testing.T) {
	rm := make(map[string]lineReport)
	rm["lineOne"] = lineReport{"lineOne", 1, []string{"fileOne"}}

	rs := getReportValues(rm)
	if rs[0].line != "lineOne" {
		t.Errorf("Expected \"lineOne\", got %s", rs[0].line)
	}
	if rs[0].count != 1 {
		t.Errorf("Expected 1, got %v", rs[0].count)
	}
}

func TestSortingReportSlice(t *testing.T) {
	rs := make(lineReports, 3)
	rs[0] = lineReport{"AlineOne", 2, []string{"fileOne"}}
	rs[1] = lineReport{"BlineOne", 2, []string{"fileOne"}}
	rs[2] = lineReport{"lineOne", 5, []string{"fileTwo"}}

	sort.Sort(rs)
	if rs[0].line != "lineOne" {
		t.Errorf("Expected \"lineOne\", got %s", rs[0].line)
	}
}

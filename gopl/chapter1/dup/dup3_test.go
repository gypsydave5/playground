package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

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

func TestUpdatingReport(t *testing.T) {
	count := 2
	filename := "file1"
	r1 := lineReport{"line", 3, []string{"file2"}}
	r2 := updateReport(r1, count, filename)
	if r2.count != 5 {
		t.Errorf("Expected 5, got %v", r2.count)
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

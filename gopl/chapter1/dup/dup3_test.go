package main

import "testing"

import "strings"

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
	rm["lineTwo"] = lineReport{"lineTwo", 2, []string{"fileTwo"}}

	ra := getReportValues(rm)
	if ra[0].line != "lineOne" {
		t.Errorf("Expected \"lineOne\", got %v", ra[0].line)
	}
}

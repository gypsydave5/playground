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
	if r["Hello"].count != 4 {
		t.Errorf("Expected 4, got %v", r["Hello"].count)
	}
}

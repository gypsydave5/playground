package main

import "testing"

func TestFToCBoiling(t *testing.T) {
	c := fToC(212)
	if c != 100 {
		t.Errorf("Expected 10, got %v", c)
	}
}

func TestsFToCFreezing(t *testing.T) {
	c := fToC(32)
	if c != 32 {
		t.Errorf("Expected 32, got %v", c)
	}
}

package main

import "testing"

func TestSquishSpace(t *testing.T) {
	b := []byte("Hello,	世界")
	b2 := squishSpace(b)

	want := "Hello, 世界"
	got := string(b2)
	if got != want {
		t.Errorf("Expected %v to equal %v", got, want)
	}

	b = []byte("Hello,	世  	  界")
	b2 = squishSpace(b)

	want = "Hello, 世 界"
	got = string(b2)
	if got != want {
		t.Errorf("Expected %v to equal %v", got, want)
	}
}

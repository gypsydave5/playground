package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchChannel(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello!")
		}))
	ch := make(chan string, 1)

	fetch(ts.URL, ch)
	out := <-ch

	if out != "hello!" {
		t.Errorf("Expected \"hello!\", but got %s", out)
	}
}

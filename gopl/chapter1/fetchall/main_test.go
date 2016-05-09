package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestFetchConcurrent(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello!")
			time.Sleep(100 * time.Millisecond)
		}))
	ch := make(chan string, 2)

	go fetchConcurrent(ts.URL, ch)
	go fetchConcurrent(ts.URL, ch)

	validReportRegex := fmt.Sprintf("^0\\.[0-9]{2}s\\s+%d\\s%s", 7, regexp.QuoteMeta(ts.URL))
	fmt.Println(validReportRegex)
	reportRegex, _ := regexp.Compile(validReportRegex)

	out := <-ch
	out2 := <-ch

	if !reportRegex.MatchString(out) {
		t.Errorf("invalid report string, %v", out)
	}

	if !reportRegex.MatchString(out2) {
		t.Errorf("invalid report string, %v", out2)
	}
}

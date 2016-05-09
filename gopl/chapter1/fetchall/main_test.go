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
	ts1 := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello!")
			time.Sleep(100 * time.Millisecond)
		}))

	ts2 := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hi")
			time.Sleep(200 * time.Millisecond)
		}))
	ch := make(chan string, 2)

	go fetchConcurrent(ts1.URL, ch)
	go fetchConcurrent(ts2.URL, ch)

	validReportRegexOne := fmt.Sprintf("^0\\.1[0-9]{1}s\\s+%d\\s%s", 7, regexp.QuoteMeta(ts1.URL))
	validReportRegexTwo := fmt.Sprintf("^0\\.2[0-9]{1}s\\s+%d\\s%s", 3, regexp.QuoteMeta(ts2.URL))
	reportRegexOne, _ := regexp.Compile(validReportRegexOne)
	reportRegexTwo, _ := regexp.Compile(validReportRegexTwo)

	out := <-ch
	out2 := <-ch

	if !reportRegexOne.MatchString(out) {
		t.Errorf("invalid report string, %v", out)
	}

	if !reportRegexTwo.MatchString(out2) {
		t.Errorf("invalid report string, %v", out2)
	}
}

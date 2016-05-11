package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type fetchReport struct {
	seconds float64
	bytes   int64
	url     string
	err     error
}

func (r *fetchReport) String() string {
	return fmt.Sprintf("%.2fs %7d %s\n", r.seconds, r.bytes, r.url)
}

func main() {
	start := time.Now()
	ch := make(chan fetchReport)

	f, err := os.Create("./results.txt")
	defer f.Close()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for _, url := range os.Args[1:] {
		go fetchConcurrent(url, ch)
	}

	for range os.Args[1:] {
		report := <-ch
		f.WriteString(report.String())
	}

	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchConcurrent(url string, ch chan<- fetchReport) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fetchReport{err: err, url: url}
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fetchReport{err: err, url: url}
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fetchReport{secs, nbytes, url, nil}
	return
}

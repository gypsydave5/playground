package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

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
		f.WriteString(<-ch)
	}

	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchConcurrent(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
	return
}

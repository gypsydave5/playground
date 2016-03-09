package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = addSchema(url)
		resp, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("<== Status Code: %v ==>\n\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func fetch(url string) (*http.Response, error) {
	return http.Get(url)
}

func fetchMany(urls []string) ([]*http.Response, error) {
	var resps []*http.Response
	var err error
	for _, url := range urls {
		resp, err := fetch(url)
		if err != nil {
			return resps, err
		}
		resps = append(resps, resp)
	}
	return resps, err
}

func addSchema(url string) string {
	newURL := url
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		newURL = "http://" + url
	}
	return newURL
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = addSchema(url)
		body, status, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("<== Status Code: %v ==>\n\n", status)
		_, err = io.Copy(os.Stdout, strings.NewReader(body))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func fetch(url string) (string, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	statusCode := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	resp.Body.Close()
	return string(body), statusCode, err
}

func fetchMany(urls []string) ([]string, error) {
	var bodies []string
	var err error
	for _, url := range urls {
		body, _, err := fetch(url)
		if err != nil {
			return bodies, err
		}
		bodies = append(bodies, body)
	}
	return bodies, err
}

func addSchema(url string) string {
	newURL := url
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		newURL = "http://" + url
	}
	return newURL
}

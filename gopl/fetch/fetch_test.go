package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPrefixAdderHttp(t *testing.T) {
	url := "http://bob.com"
	result := addSchema(url)
	if result != url {
		t.Errorf(`expected %q, got %q`, url, result)
	}
}

func TestAddSchemaHttps(t *testing.T) {
	url := "https://bob.com"
	result := addSchema(url)
	if result != url {
		t.Errorf(`expected %q, got %q`, url, result)
	}
}
func TestAddSchemaNoSchema(t *testing.T) {
	url := "bob.com"
	result := addSchema(url)
	expectation := "http://bob.com"
	if result != expectation {
		t.Errorf(`expected %q, got %q`, expectation, result)
	}
}

func TestMakingGoodHTTP(t *testing.T) {
	expectation := "Hello, delicious friend"
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, expectation)
		}))
	defer ts.Close()

	resp, _ := fetch(ts.URL)
	b, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if string(b) != expectation {
		t.Errorf(`expected %q, got %q`, expectation, string(b))
	}
}

func TestMakingBadHTTP(t *testing.T) {
	_, err := fetch("http://badbadurl")
	if err == nil {
		t.Errorf(`expected fetch to error when there's no server`)
	}
}

func TestFetchingMany(t *testing.T) {
	callCount := 0
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			callCount++
			responseString := "Call number " + strconv.Itoa(callCount)
			fmt.Fprint(w, responseString)
		}))
	defer ts.Close()

	urls := []string{ts.URL + "/first", ts.URL + "/second"}
	resps, _ := fetchMany(urls)
	b1, _ := ioutil.ReadAll(resps[0].Body)
	b2, _ := ioutil.ReadAll(resps[1].Body)
	if string(b1) != "Call number 1" {
		t.Errorf(`expected %q, got %q`, "Call number 1", string(b1))
	}
	if string(b2) != "Call number 2" {
		t.Errorf(`expected %q, got %q`, "Call number 2", string(b2))
	}
}

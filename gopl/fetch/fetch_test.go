package main

import (
	"fmt"
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

	b, _, _ := fetch(ts.URL)

	if string(b) != expectation {
		t.Errorf(`expected %q, got %q`, expectation, string(b))
	}
}

func TestMakingBadHTTP(t *testing.T) {
	_, _, err := fetch("http://badbadurl")
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
	bodies, _ := fetchMany(urls)
	if bodies[0] != "Call number 1" {
		t.Errorf(`expected %q, got %q`, "Call number 1", bodies[0])
	}
	if bodies[1] != "Call number 2" {
		t.Errorf(`expected %q, got %q`, "Call number 2", bodies[1])
	}
}

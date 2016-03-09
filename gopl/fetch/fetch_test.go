package main

import "testing"

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

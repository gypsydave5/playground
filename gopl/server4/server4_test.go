package main

import (
	"net/url"
	"testing"
)

func TestLissajousOptsCycles(t *testing.T) {
	form := map[string][]string{
		"cycles": []string{"20"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Cycles != 20.0 {
		t.Errorf("expected Cycles to be 20.0, got %v", opts.Cycles)
	}
}

func TestLissajousOptsCyclesDefault(t *testing.T) {
	form := map[string][]string{}
	opts := lissajousOpts(url.Values(form))
	if opts.Cycles != 5.0 {
		t.Errorf("expected Cycles to be 5.0, got %v", opts.Cycles)
	}
}

func TestLissajousOptsResolution(t *testing.T) {
	form := map[string][]string{
		"resolution": []string{"0.05"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Resolution != 0.05 {
		t.Errorf("expected Resolution to be 0.05, got %v", opts.Resolution)
	}
}

func TestLissajousOptsResolutionDefault(t *testing.T) {
	form := map[string][]string{}
	opts := lissajousOpts(url.Values(form))
	if opts.Resolution != 0.001 {
		t.Errorf("expected Resolution to be 0.001, got %v", opts.Resolution)
	}
}

func TestLissajousOptsSize(t *testing.T) {
	form := map[string][]string{
		"size": []string{"200"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Size != 200 {
		t.Errorf("expected Size to be 200, got %v", opts.Size)
	}
}

func TestLissajousOptsSizeDefault(t *testing.T) {
	form := map[string][]string{}
	opts := lissajousOpts(url.Values(form))
	if opts.Size != 100 {
		t.Errorf("expected Size to be 100, got %v", opts.Size)
	}
}

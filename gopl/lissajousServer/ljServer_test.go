package main

import (
	"fmt"
	"net/url"
	"testing"
	"testing/quick"
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

func TestLissajousOptsFrammes(t *testing.T) {
	form := map[string][]string{
		"frames": []string{"128"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Frames != 128 {
		t.Errorf("expected Frames to be 200, got %v", opts.Frames)
	}
}

func TestLissajousOptsFramesDefault(t *testing.T) {
	form := map[string][]string{}
	opts := lissajousOpts(url.Values(form))
	if opts.Frames != 64 {
		t.Errorf("expected Frames to be 64, got %v", opts.Size)
	}
}

func TestLissajousOptsDelay(t *testing.T) {
	form := map[string][]string{
		"delay": []string{"286"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Delay != 286 {
		t.Errorf("expected Delay to be 286, got %v", opts.Frames)
	}
}

func TestLissajousOptsDelayDefault(t *testing.T) {
	form := map[string][]string{}
	opts := lissajousOpts(url.Values(form))
	if opts.Delay != 8 {
		t.Errorf("expected Delay to be 8, got %v", opts.Size)
	}
}

func TestLissajousOptsFreqency(t *testing.T) {
	form := map[string][]string{
		"frequency": []string{"1.337"},
	}
	opts := lissajousOpts(url.Values(form))
	if opts.Frequency != 1.337 {
		t.Errorf("expected Frequency to be 1.337, got %v", opts.Frequency)
	}
}

// Default Frequncy is always between 0.0 and 3.0
func TestLissaljousFrequencyDefault(t *testing.T) {
	f := func() bool {
		form := map[string][]string{}
		opts := lissajousOpts(url.Values(form))
		fmt.Printf("trying: %v\n", opts)
		return opts.Frequency >= 0 && opts.Frequency <= 3
	}
	if err := quick.Check(f, nil); err != nil {
		t.Errorf("expected default Frequncey to be between 0 and 3: %v", err)
	}
}

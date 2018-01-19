package main

import (
	"net/url"
	"testing"

	"github.com/gypsydave5/playground/gopl/chapter3/fractal"
)

func TestUrlParamsBounds(t *testing.T) {
	params := map[string][]string{"x": {"0.1"}, "y": {"0.2"}, "zoom": {"2.0"}}
	urlParams := url.Values(params)
	opts := applyOptions(fractal.Parameters{}, urlParams)

	if opts.Bounds.Xmax != 4.1 {
		t.Error("Expected Xmax to be 4.1, but got: ", opts.Bounds.Xmax)
	}

	if opts.Bounds.Xmin != -3.9 {
		t.Error("Expected Xmin to be 3.9, but got: ", opts.Bounds.Xmin)
	}

	if opts.Bounds.Ymax != 4.2 {
		t.Error("Expected Ymax to be 4.1, but got: ", opts.Bounds.Ymax)
	}

	if opts.Bounds.Ymin != -3.8 {
		t.Error("Expected Ymin to be 3.9, but got: ", opts.Bounds.Ymin)
	}
}

func TestUrlParamsIterations(t *testing.T) {
	params := map[string][]string{"iterations": {"255"}}

	urlParams := url.Values(params)
	opts := applyOptions(fractal.Parameters{}, urlParams)

	if opts.Iterations != 255 {
		t.Error("Expected Iterations to be 255, but got: ", opts.Iterations)
	}
}

func TestUrlParamsWidthHeight(t *testing.T) {
	params := map[string][]string{"width": {"1337"}, "height": {"808"}}

	urlParams := url.Values(params)
	opts := applyOptions(fractal.Parameters{}, urlParams)

	if opts.Width != 1337 {
		t.Error("Expected Width to be 1337, but got: ", opts.Width)
	}

	if opts.Height != 808 {
		t.Error("Expected Height to be 808, but got: ", opts.Height)
	}
}

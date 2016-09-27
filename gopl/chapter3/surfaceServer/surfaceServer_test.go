package main

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gypsydave5/playground/gopl/chapter3/surface"
)

func TestHandlerContentType(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, r)

	contentType := w.Header().Get("Content-Type")
	if contentType != "image/svg+xml" {
		t.Error("Expected xml Content-Type, but got", contentType)
	}
}

func TestUrlParamsCells(t *testing.T) {
	params := map[string][]string{"cells": {"500"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)

	if opts.Cells != 500 {
		t.Error("Expected cells to be 500, but got: ", opts.Cells)
	}
}

func TestUrlParamsWidth(t *testing.T) {
	params := map[string][]string{"width": {"500"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)

	if opts.Width != 500 {
		t.Error("Expected width to be 500, but got: ", opts.Width)
	}
}

func TestUrlParamsHeight(t *testing.T) {
	params := map[string][]string{"height": {"500"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)

	if opts.Height != 500 {
		t.Error("Expected height to be 500, but got: ", opts.Height)
	}
}

func TestUrlParamsXYRange(t *testing.T) {
	params := map[string][]string{"xyrange": {"45.0"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)

	if opts.XYRange != 45.0 {
		t.Error("Expected XYRange to be 45.0, but got: ", opts.XYRange)
	}
}

func TestUrlParamsLowerColor(t *testing.T) {
	params := map[string][]string{"lowercolor": {"bada55"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)

	if opts.LowerColor != "bada55" {
		t.Error("Expected LoweColor to be 'bada55', but got: ", opts.LowerColor)
	}
}

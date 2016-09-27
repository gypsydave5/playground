package main

import (
	"image/color"
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
	params := map[string][]string{"lowercolor": {"ff00ff"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)
	expectedRGBA := color.RGBA{255, 0, 255, 0}

	if opts.LowerColor != expectedRGBA {
		t.Error("Expected LoweColor to be 'bada55', but got: ", opts.LowerColor)
	}
}

func TestUrlParamsLowerColorError(t *testing.T) {
	params := map[string][]string{"lowercolor": {"PUNKER"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)
	expectedRGBA := color.RGBA{0, 0, 255, 0}

	if opts.LowerColor != expectedRGBA {
		t.Error("Expected LoweColor to be: ", expectedRGBA, " but got: ", opts.LowerColor)
	}
}

func TestRGBAfromHex(t *testing.T) {
	hex := "FF7F01"

	rgba, _ := rgbaFromHex(hex)

	expectedRGBA := color.RGBA{255, 127, 1, 0}
	if rgba != expectedRGBA {
		t.Error("Expected: ", expectedRGBA, " but got: ", rgba)
	}
}

func TestRGBAfromHexError(t *testing.T) {
	hex := "BOOMER"

	_, err := rgbaFromHex(hex)

	if err == nil {
		t.Error("Expected: an error,  but got: ", err)
	}
}

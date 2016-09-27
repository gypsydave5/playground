package main

import (
	"image/color"
	"net/url"
	"testing"

	"github.com/gypsydave5/playground/gopl/chapter3/surface"
)

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
		t.Error("Expected LowerColor to be 'bada55', but got: ", opts.LowerColor)
	}
}

func TestUrlParamsLowerColorError(t *testing.T) {
	params := map[string][]string{"lowercolor": {"PUNKER"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)
	expectedRGBA := color.RGBA{0, 0, 255, 0}

	if opts.LowerColor != expectedRGBA {
		t.Error("Expected LowerColor to be: ", expectedRGBA, " but got: ", opts.LowerColor)
	}
}

func TestUrlParamsUpperColor(t *testing.T) {
	params := map[string][]string{"uppercolor": {"7f7f7f"}}
	urlParams := url.Values(params)
	opts := applyOptions(surface.NewOptions(), urlParams)
	expectedRGBA := color.RGBA{127, 127, 127, 0}

	if opts.UpperColor != expectedRGBA {
		t.Error("Expected UpperColor to be: ", expectedRGBA, " but got: ", opts.UpperColor)
	}
}

package main

import (
	"image/color"
	"testing"
)

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

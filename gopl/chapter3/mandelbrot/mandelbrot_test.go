package main

import (
	"image/color"
	"testing"
)

const iterations = 200
const contrast = 15

func TestMandleBrotOrigin(t *testing.T) {
	num := complex(0, 0)

	expectedColor := color.Black

	c := mandelbrotShade(num, iterations, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

func TestMandleBrotEdge(t *testing.T) {
	num := complex(-1, 0.4)

	expectedColor := color.Gray{165}

	c := mandelbrotShade(num, iterations, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

func TestMandleBrotNegativeEdge(t *testing.T) {
	num := complex(-1, -0.4)

	expectedColor := color.Gray{165}

	c := mandelbrotShade(num, iterations, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

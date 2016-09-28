package main

import (
	"image/color"
	"testing"
)

func TestMandleBrotOrigin(t *testing.T) {
	const iterations = 200
	const contrast = 15
	num := complex(0, 0)

	expectedColor := color.Black

	tries, escaped := escapeIteration(num, iterations)
	c := mandelbrotShade(tries, escaped, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

func TestMandleBrotEdge(t *testing.T) {
	const iterations = 200
	const contrast = 15
	num := complex(-1, 0.4)

	expectedColor := color.Gray{165}

	tries, escaped := escapeIteration(num, iterations)
	c := mandelbrotShade(tries, escaped, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

func TestMandleBrotNegativeEdge(t *testing.T) {
	const iterations = 200
	const contrast = 15
	num := complex(-1, -0.4)

	expectedColor := color.Gray{165}

	tries, escaped := escapeIteration(num, iterations)
	c := mandelbrotShade(tries, escaped, contrast)

	if c != expectedColor {
		t.Error("Expected ", expectedColor, ", but got: ", c)
	}
}

package main

import (
	"image/color"
	"math"
	"testing"
	"testing/quick"
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

func TestEscapeIterationZeroNeverEscapes(t *testing.T) {
	f := func(iterations uint8) bool {
		num := complex(0, 0)
		_, escaped := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("0i0 should always be in Mandelbrot set")
	}

}

func TestEscapeIterationNegativeRealNeverEscapes(t *testing.T) {
	f := func(iterations uint8, real float64) bool {
		fromNeg2toPoint25 := (math.Abs(math.Mod(real, 2.2499999))) - 2

		num := complex(fromNeg2toPoint25, 0)
		_, escaped := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("complex numbers with a negative real component and no complex component should always be in the Mnadelbrot set")
	}

}

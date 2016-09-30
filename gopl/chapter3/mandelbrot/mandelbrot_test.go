package main

import (
	"image/color"
	"math"
	"testing"
	"testing/quick"
)

func TestEscapeIterationZeroNeverEscapes(t *testing.T) {
	f := func(iterations uint8) bool {
		num := complex(0, 0)
		_, escaped, _ := escapeIteration(num, iterations)
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
		_, escaped, _ := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("complex numbers with a negative real component and no complex component should always be in the Mnadelbrot set")
	}
}

func TestHSVColorBlack(t *testing.T) {
	h := HSVA{0, 0, 0, 1}

	r, g, b, _ := h.RGBA()
	ar, ag, ab, _ := color.Black.RGBA()

	if r != ar {
		t.Error("Expected : ", ar, "but got : ", r)
	}

	if g != ab {
		t.Error("Expected : ", ag, "but got : ", g)
	}

	if b != ag {
		t.Error("Expected : ", ab, "but got : ", b)
	}
}

func TestHSVColorWhite(t *testing.T) {
	h := HSVA{0, 0, 1, 1}

	r, g, b, _ := h.RGBA()
	ar, ag, ab, _ := color.White.RGBA()

	if r != ar {
		t.Error("Expected : ", ar, "but got : ", r)
	}

	if g != ab {
		t.Error("Expected : ", ag, "but got : ", g)
	}

	if b != ag {
		t.Error("Expected : ", ab, "but got : ", b)
	}
}

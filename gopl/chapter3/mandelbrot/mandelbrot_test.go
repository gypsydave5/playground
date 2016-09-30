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

func TestHSVColors(t *testing.T) {
	hsvas := []HSVA{
		HSVA{0, 0, 0, 1},
		HSVA{0, 0, 1, 1},
		HSVA{0, 1, 1, 1},
	}
	colors := []color.Color{
		color.Black,
		color.White,
		color.RGBA{0xff, 0, 0, 0xff},
	}

	for i, h := range hsvas {
		hr, hb, hg, ha := h.RGBA()
		cr, cb, cg, ca := colors[i].RGBA()

		if hr != cr {
			t.Error("Expected red : ", cr, "but got : ", hr)
		}

		if hg != cb {
			t.Error("Expected green : ", cg, "but got : ", hg)
		}

		if hb != cg {
			t.Error("Expected blue : ", cb, "but got : ", hb)
		}

		if ha != ca {
			t.Error("Expected alpha : ", ca, "but got : ", ha)
		}

	}
}

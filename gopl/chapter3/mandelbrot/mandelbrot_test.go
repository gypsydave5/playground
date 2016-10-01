package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type neg2ToZeroPt25 float64

func (m neg2ToZeroPt25) Generate(rand *rand.Rand, size int) reflect.Value {
	mn := neg2ToZeroPt25((rand.Float64() * 2.25) - 2)
	return reflect.ValueOf(mn)
}

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
	f := func(iterations uint8, real neg2ToZeroPt25) bool {

		fmt.Println(real)
		num := complex(real, 0)
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
		HSVA{0, 1, 1, 1},   //Red
		HSVA{120, 1, 1, 1}, //Green
		HSVA{240, 1, 1, 1}, //Blue
		HSVA{60, 1, 1, 1},  //Yellow
		HSVA{180, 1, 1, 1}, //Cyan
		HSVA{300, 1, 1, 1}, //Magenta
	}
	colors := []color.Color{
		color.Black,
		color.White,
		color.RGBA{0xff, 0, 0, 0xff},    //Red
		color.RGBA{0, 0xff, 0, 0xff},    //Green
		color.RGBA{0, 0, 0xff, 0xff},    //Blue
		color.RGBA{0xff, 0xff, 0, 0xff}, //Yellow
		color.RGBA{0, 0xff, 0xff, 0xff}, //Cyan
		color.RGBA{0xff, 0, 0xff, 0xff}, //Magenta
	}

	for i, h := range hsvas {
		hr, hb, hg, ha := h.RGBA()
		cr, cb, cg, ca := colors[i].RGBA()

		if hr != cr {
			t.Errorf("Expected red : %v, but got %v [HSVA%v]", cr, hr, h)
		}

		if hg != cg {
			t.Errorf("Expected green : %v, but got %v [HSVA%v]", cg, hg, h)
		}

		if hb != cb {
			t.Errorf("Expected blue : %v, but got %v [HSVA%v]", cb, hb, h)
		}

		if ha != ca {
			t.Errorf("Expected alpha : %v, but got %v [HSVA%v]", ca, ha, h)
		}

	}
}

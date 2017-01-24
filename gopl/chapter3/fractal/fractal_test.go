package fractal

import (
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
		num := complex(real, 0)
		_, escaped, _ := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("complex numbers with a negative real component and no complex " +
			"component should always be in the Mandelbrot set")
	}
}

func TestAverageColor(t *testing.T) {
	color1 := color.Gray{0}
	color2 := color.Gray{254}

	avgColor := averageColor(color1, color2)
	r, g, b, a := color.Gray{127}.RGBA()
	expectedColor := color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}

	if avgColor != expectedColor {
		t.Error("Was expecting", expectedColor, "and yet we received", avgColor, ". How disappointing.")
	}
}

func TestPixelToCoord(t *testing.T) {
	vp := vpixel{100, 100}
	bounds := Bounds{
		Xmax: 100,
		Xmin: -100,
		Ymax: 100,
		Ymin: -100,
	}

	params := MandelbrotParameters{
		Bounds: bounds,
		Width:  200,
		Height: 200,
	}

	gc := pixelToCoord(vp, params)
	expected := Coord{0, 0}

	if expected != gc {
		t.Error("Was expecting", expected, "and yet we received", gc, ". How disappointing.")
	}
}

func TestCoordToComplex(t *testing.T) {
	gc := Coord{1, 1}
	z := coordToComplex(gc)

	if z != complex(1, 1) {
		t.Error("Expected", complex(1, 1), "but got", z)
	}
}

func TestCoordsZoomToBounds(t *testing.T) {
	center := Coord{1, 0.5}
	zoom := 0.5

	var b = *CoordsZoomToBounds(center, zoom, 2)

	expected := Bounds{0, -0.5, 2, 1.5}

	if b != expected {
		t.Error("Expected", expected, "but got", b)
	}
}

func TestCoordsZoomToBoundsMidCoordsAndZoomOneActsAsIdentity(t *testing.T) {
	center := Coord{0, 0}
	zoom := 1.0

	f := func(defaultBound float64) bool {
		var b = CoordsZoomToBounds(center, zoom, defaultBound)
		return (b.Xmax == defaultBound) &&
			(b.Xmin == -defaultBound) &&
			(b.Ymax == defaultBound) &&
			(b.Ymin == -defaultBound)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("coordsZoomToBounds not performing as identity")
	}
}

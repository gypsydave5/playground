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

func TestSuperSample(t *testing.T) {
	vp := vpixel{0.0, 0.0}
	params := MandelbrotParameters{
		Width:  200,
		Height: 200,
		Xmax:   100,
		Xmin:   -100,
		Ymax:   100,
		Ymin:   -100,
	}
	pcFun := newMandelbrotPixelColorFunction(10, params, alwaysWhiteShader)
	colors := superSample(vp, pcFun)
	expectedColors := []color.Color{
		color.White,
		color.White,
		color.White,
		color.White,
	}

	if !reflect.DeepEqual(colors, expectedColors) {
		t.Error("Was expecting", expectedColors, "and yet we received", colors, ". How disappointing.")
	}
}

func TestPixelToCoord(t *testing.T) {
	vp := vpixel{100, 100}
	params := MandelbrotParameters{
		Width:  200,
		Height: 200,
		Xmax:   100,
		Xmin:   -100,
		Ymax:   100,
		Ymin:   -100,
	}

	gc := pixelToCoord(vp, params)
	expected := coord{0, 0}

	if expected != gc {
		t.Error("Was expecting", expected, "and yet we received", gc, ". How disappointing.")
	}
}

func TestCoordToComplex(t *testing.T) {
	gc := coord{1, 1}
	z := coordToComplex(gc)

	if z != complex(1, 1) {
		t.Error("Expected", complex(1, 1), "but got", z)
	}
}

func alwaysWhiteShader(tries, maxTries uint8, escaped bool, contrast int, zFinal complex128) color.Color {
	return color.White
}

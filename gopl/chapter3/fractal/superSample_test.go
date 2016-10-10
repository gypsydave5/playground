package fractal

import (
	"image/color"
	"reflect"
	"testing"
)

func TestSuperSample(t *testing.T) {
	vp := vpixel{0.0, 0.0}

	c := superSample(vp, alwaysWhitePixel)
	expectedColor := color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff}

	if !reflect.DeepEqual(c, expectedColor) {
		t.Errorf("Was expecting %#v, and yet we received %#v. How disappointing.", expectedColor, c)
	}
}

func alwaysWhitePixel(p vpixel) color.Color {
	return color.White
}

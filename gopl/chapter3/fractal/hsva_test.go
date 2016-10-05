package fractal

import (
	"image/color"
	"testing"
)

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
		c := colors[i]
		hr, hb, hg, ha := h.RGBA()
		cr, cb, cg, ca := c.RGBA()

		if hr != cr {
			t.Errorf("Expected red : %v, but got %v [HSVA%v] [RGBA%v]", cr, hr, h, c)
		}

		if hg != cg {
			t.Errorf("Expected green : %v, but got %v [HSVA%v] [RGBA%v]", cg, hg, h, c)
		}

		if hb != cb {
			t.Errorf("Expected blue : %v, but got %v [HSVA%v] [RGBA%v]", cb, hb, h, c)
		}

		if ha != ca {
			t.Errorf("Expected alpha : %v, but got %v [HSVA%v] [RGBA%v]", ca, ha, h, c)
		}
	}
}

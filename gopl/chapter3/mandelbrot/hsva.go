package main

import "math"

// HSVA represents colour by hue, saturation, value and alpha.
// Hue should be between 0 and 360
// Saturation, value and alpha between 0 and 1
type HSVA struct {
	H float64
	S float64
	V float64
	A float64
}

// RGBA returns the HSVA color as RGBA values as per the color.Color interface
func (h HSVA) RGBA() (r uint32, g uint32, b uint32, a uint32) {
	var rf, gf, bf, af float64

	c := h.S * h.V
	x := c * (1 - math.Abs(math.Mod(h.H/60, 2.0)-1.0))
	m := h.V - c

	switch {
	case h.H < 60:
		rf = c
		gf = x
		bf = 0
		af = h.A
	case h.H < 120:
		rf = x
		gf = c
		bf = 0
		af = h.A
	case h.H < 180:
		rf = 0
		gf = c
		bf = x
		af = h.A
	case h.H < 240:
		rf = 0
		gf = x
		bf = c
		af = h.A
	case h.H < 300:
		rf = x
		gf = 0
		bf = c
		af = h.A
	case h.H < 360:
		rf = c
		gf = 0
		bf = x
		af = h.A
	}

	r = uint32((rf + m) * 0xffff)
	g = uint32((gf + m) * 0xffff)
	b = uint32((bf + m) * 0xffff)
	a = uint32(af * 0xffff)

	return r, g, b, a
}

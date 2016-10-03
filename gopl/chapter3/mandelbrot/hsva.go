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
	var rf, gf, bf float64
	hp := h.H / 60.0

	c := h.S * h.V
	x := c * (1.0 - math.Abs(math.Mod(hp, 2.0)-1.0))
	m := h.V - c

	switch {
	case h.H < 60:
		rf = c
		gf = x
	case h.H < 120:
		rf = x
		gf = c
	case h.H < 180:
		gf = c
		bf = x
	case h.H < 240:
		gf = x
		bf = c
	case h.H < 300:
		rf = x
		bf = c
	case h.H < 360:
		rf = c
		bf = x
	}

	r = uint32((rf + m) * 65535.0)
	g = uint32((gf + m) * 65535.0)
	b = uint32((bf + m) * 65535.0)
	a = uint32(h.A * 65535.0)

	return r, g, b, a
}

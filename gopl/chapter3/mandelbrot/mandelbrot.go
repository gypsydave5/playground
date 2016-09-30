// Mandelbrot emits an animated GIF of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/cmplx"
	"os"

	"github.com/andybons/gogif"
)

// HSVA represents colour by hue, saturation, value and alpha.
// Hue should be between 0 and 360
// Saturation, value and alpha between 0 and 1
type HSVA struct {
	H float64
	S float64
	V float64
	A float64
}

func (h HSVA) RGBA() (r uint32, g uint32, b uint32, a uint32) {
	var rf, gf, bf, af float64

	c := h.S * h.V
	x := c * (1 - math.Abs(math.Mod(h.H/60, 2.0)-1.0))
	m := h.V - c

	rf = c
	gf = x
	bf = 0
	af = h.A

	r = uint32((rf + m) * 0xffff)
	g = uint32((gf + m) * 0xffff)
	b = uint32((bf + m) * 0xffff)
	a = uint32(af * 0xffff)
	return r, g, b, a
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 60
	contrast               = 15
	delay                  = 20
)

func main() {
	anim := gif.GIF{LoopCount: iterations}

	animateMandelbrot(iterations, width, height, &anim)

	gif.EncodeAll(os.Stdout, &anim)
}

func animateMandelbrot(maxIterations uint8, width int, height int, anim *gif.GIF) {
	for i := uint8(0); i < iterations; i++ {
		img := generateMandelbrot(i, width, height)
		palettedImage := rgbaToPalleted(img)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, palettedImage)
	}
}

func generateMandelbrot(iterations uint8, width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)

			tries, escaped, _ := escapeIteration(z, iterations)

			shade := mandelbrotShade(tries, escaped, contrast)
			img.Set(px, py, shade)
		}
	}
	return img
}

func rgbaToPalleted(rgba *image.RGBA) *image.Paletted {
	bounds := rgba.Bounds()
	palettedImage := image.NewPaletted(bounds, nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	quantizer.Quantize(palettedImage, bounds, rgba, image.ZP)
	return palettedImage
}

func mandelbrotShade(tries uint8, escaped bool, contrast int) color.Color {
	if !escaped {
		return color.Black
	}

	return color.Gray{255 - uint8(contrast)*tries}
}


func escapeIteration(z complex128, maxIterations uint8) (iterations uint8, escaped bool, zFinal complex128) {
	var v complex128

	for iterations = uint8(0); iterations < maxIterations; iterations++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return iterations, true, v
		}
	}

	return iterations, false, v
}

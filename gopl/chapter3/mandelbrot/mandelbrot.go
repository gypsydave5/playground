// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"math/cmplx"
	"os"

	"github.com/andybons/gogif"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 10
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

			// image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotShade(z, iterations, contrast))
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

func mandelbrotShade(z complex128, iterations uint8, contrast uint8) color.Color {
	tries, escaped := escapeIteration(z, iterations)

	if !escaped {
		return color.Black
	}

	return color.Gray{255 - contrast*tries}
}

func escapeIteration(z complex128, maxIterations uint8) (iterations uint8, escaped bool) {
	var v complex128

	for iterations = uint8(0); iterations < maxIterations; iterations++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return iterations, true
		}
	}

	return iterations, false
}

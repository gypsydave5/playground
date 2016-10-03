// Mandelbrot emits an animated GIF of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/cmplx"
	"os"

	"github.com/andybons/gogif"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 256, 256
	maxIterations          = 10
	startingIteration      = 5
	contrast               = 15
	delay                  = 20
)

func main() {
	anim := gif.GIF{LoopCount: maxIterations}
	animateMandelbrot(maxIterations, startingIteration, width, height, &anim)
	gif.EncodeAll(os.Stdout, &anim)
}

func animateMandelbrot(maxIterations, startingIteration uint8, width, height int, anim *gif.GIF) {
	for i := uint8(startingIteration); i < maxIterations; i++ {
		img := generateMandelbrot(i, width, height)
		palettedImage := rgbaToPalleted(img)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, palettedImage)
	}
}

func generateMandelbrot(iterations uint8, width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)

			tries, escaped, zFinal := escapeIteration(z, iterations)

			shade := colorShade(tries, iterations, escaped, contrast, zFinal)
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

func greyShade(tries uint8, escaped bool, contrast int, zFinal complex128) color.Color {
	if !escaped {
		return color.Black
	}

	return color.Gray{255 - uint8(contrast)*tries}
}

func colorShade(tries, maxTries uint8, escaped bool, contrast int, zFinal complex128) color.Color {
	if !escaped {
		return color.Black
	}

	return smoothHSV(tries, maxTries, zFinal)
}

func escapeIteration(z complex128, mi uint8) (i uint8, escaped bool, zFinal complex128) {
	var v complex128
	for i = uint8(0); i < mi; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return i, true, v
		}
	}

	return i, false, v
}

func smoothHSV(tries, maxTries uint8, zFinal complex128) HSVA {
	hue := math.Abs((float64(tries) / float64(maxTries)) * 360)
	log.Println("Hue : ", hue)

	return HSVA{
		H: hue,
		S: 0.8,
		V: 1.0,
		A: 1.0,
	}
}

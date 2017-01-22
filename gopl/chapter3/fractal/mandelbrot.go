// Package fractal presents functions to write fractal images to various image filetypes.
// It currently supports images in the Mandelbrot set as both PNG images and animated GIFs
package fractal

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"math/cmplx"

	"github.com/andybons/gogif"
)

var loggingEnabled bool

// MandelbrotParameters supplies parameters for the generation of a Mandelbrot
// image
type MandelbrotParameters struct {
	Xmin, Ymin, Xmax, Ymax         float64
	Width, Height, Contrast, Delay int
	Iterations, StartingIteration  uint8
	Logging, Colour, SuperSample   bool
}

type coord struct {
	X float64
	Y float64
}

// WritePNG writes the Mandelbrot image to a provided io.Writer, encoded as a PNG,
// using the supplied MandelbrotParameters for configuration
func WritePNG(w io.Writer, params MandelbrotParameters) {
	loggingEnabled = params.Logging
	img := generateMandelbrot(params.Iterations, params)
	png.Encode(w, img)
}

// WriteAnimatedGIF writes an animation of the Mandelbrot fractal, increasing
// the number of iterations per frame
func WriteAnimatedGIF(w io.Writer, params MandelbrotParameters) {
	loggingEnabled = params.Logging
	anim := gif.GIF{LoopCount: int(params.Iterations)}
	animateMandelbrot(&anim, params)
	gif.EncodeAll(w, &anim)
}

func animateMandelbrot(anim *gif.GIF, params MandelbrotParameters) {
	for i := params.StartingIteration; i < params.Iterations; i++ {
		img := generateMandelbrot(i, params)
		palettedImage := rgbaToPalleted(img)
		anim.Delay = append(anim.Delay, params.Delay)
		anim.Image = append(anim.Image, palettedImage)
	}
}

func generateMandelbrot(iterations uint8, params MandelbrotParameters) *image.NRGBA {
	var mdbFn pixelToColor
	if params.Colour == true {
		mdbFn = newMandelbrotPixelColorFunction(iterations, params, colorShade)
	} else {
		mdbFn = newMandelbrotPixelColorFunction(iterations, params, greyShade)
	}

	img := image.NewNRGBA(image.Rect(0, 0, params.Width, params.Height))

	for py := 0; py < params.Height; py++ {
		for px := 0; px < params.Width; px++ {
			vp := vpixel{float64(px), float64(py)}
			var color color.Color

			if params.SuperSample {
				color = superSample(vp, mdbFn)
			} else {
				color = mdbFn(vp)
			}
			img.Set(px, py, color)
		}
	}
	return img
}

func newMandelbrotPixelColorFunction(iterations uint8, params MandelbrotParameters, sh shader) pixelToColor {
	return func(p vpixel) color.Color {
		c := pixelToCoord(p, params)
		z := coordToComplex(c)
		return complexToMandelbrotColor(z, iterations, params, sh)
	}
}

func rgbaToPalleted(rgba image.Image) *image.Paletted {
	bounds := rgba.Bounds()
	palettedImage := image.NewPaletted(bounds, nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	quantizer.Quantize(palettedImage, bounds, rgba, image.ZP)
	return palettedImage
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

func pixelToCoord(vp vpixel, params MandelbrotParameters) coord {
	x := float64(vp.X)/float64(params.Width)*(params.Xmax-params.Xmin) + (params.Xmin)
	y := float64(vp.Y)/float64(params.Height)*(params.Ymax-params.Ymin) + (params.Ymin)
	return coord{x, y}
}

func coordToComplex(gc coord) complex128 {
	return complex(gc.X, gc.Y)
}

func complexToMandelbrotColor(z complex128, iterations uint8, params MandelbrotParameters, sh shader) color.Color {
	t, e, zf := escapeIteration(z, iterations)
	return sh(t, iterations, e, params.Contrast, zf)
}

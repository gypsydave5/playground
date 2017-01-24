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
	Bounds                         Bounds
	Width, Height, Contrast, Delay int
	Iterations, StartingIteration  uint8
	Logging, Colour, SuperSample   bool
}

// Bounds represents the maximum and minimum x and y Cartesian coordinates
// for a fractal image.
type Bounds struct {
	Xmin, Ymin, Xmax, Ymax float64
}

// Coord represents a Cartesian coordinate.
type Coord struct {
	X float64
	Y float64
}

// CoordsZoomToBounds calculates the boundaries of an image based upon a central
// point, a zoom factor, and a default starting boundary size.
func CoordsZoomToBounds(centre Coord, zoom float64, defaultBound float64) *Bounds {
	return &Bounds{
		Xmax: centre.X + (zoom * defaultBound),
		Ymax: centre.Y + (zoom * defaultBound),
		Xmin: centre.X - (zoom * defaultBound),
		Ymin: centre.Y - (zoom * defaultBound),
	}
}

// WritePNG writes the Mandelbrot image to a provided io.Writer, encoded as a PNG,
// using the supplied MandelbrotParameters for configuration
func WritePNG(w io.Writer, p MandelbrotParameters) {
	loggingEnabled = p.Logging
	img := generateMandelbrot(p.Iterations, p)
	png.Encode(w, img)
}

// WriteAnimatedGIF writes an animation of the Mandelbrot fractal, increasing
// the number of iterations per frame
func WriteAnimatedGIF(w io.Writer, p MandelbrotParameters) {
	loggingEnabled = p.Logging
	anim := gif.GIF{LoopCount: int(p.Iterations)}
	animateMandelbrot(&anim, p)
	gif.EncodeAll(w, &anim)
}

func animateMandelbrot(anim *gif.GIF, p MandelbrotParameters) {
	for i := p.StartingIteration; i < p.Iterations; i++ {
		img := generateMandelbrot(i, p)
		palettedImage := rgbaToPalleted(img)
		anim.Delay = append(anim.Delay, p.Delay)
		anim.Image = append(anim.Image, palettedImage)
	}
}

func generateMandelbrot(iterations uint8, p MandelbrotParameters) *image.NRGBA {
	var mdbFn pixelToColor

	if p.Colour == true {
		mdbFn = newMandelbrotPixelColorFunction(iterations, p, colorShade)
	} else {
		mdbFn = newMandelbrotPixelColorFunction(iterations, p, greyShade)
	}

	img := image.NewNRGBA(image.Rect(0, 0, p.Width, p.Height))

	for py := 0; py < p.Height; py++ {
		for px := 0; px < p.Width; px++ {
			vp := vpixel{float64(px), float64(py)}
			var color color.Color

			if p.SuperSample {
				color = superSample(vp, mdbFn)
			} else {
				color = mdbFn(vp)
			}
			img.Set(px, py, color)
		}
	}
	return img
}

func newMandelbrotPixelColorFunction(iterations uint8, p MandelbrotParameters, sh shader) pixelToColor {
	return func(vp vpixel) color.Color {
		c := pixelToCoord(vp, p)
		z := coordToComplex(c)
		return complexToMandelbrotColor(z, iterations, p, sh)
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

func pixelToCoord(vp vpixel, p MandelbrotParameters) Coord {
	x := float64(vp.X)/float64(p.Width)*(p.Bounds.Xmax-p.Bounds.Xmin) + (p.Bounds.Xmin)
	y := float64(vp.Y)/float64(p.Height)*(p.Bounds.Ymax-p.Bounds.Ymin) + (p.Bounds.Ymin)
	return Coord{x, y}
}

func coordToComplex(gc Coord) complex128 {
	return complex(gc.X, gc.Y)
}

func complexToMandelbrotColor(z complex128, iterations uint8, p MandelbrotParameters, sh shader) color.Color {
	t, e, zf := escapeIteration(z, iterations)
	return sh(t, iterations, e, p.Contrast, zf)
}

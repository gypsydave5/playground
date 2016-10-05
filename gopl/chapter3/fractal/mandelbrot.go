// Package fractal emits an animated GIF of the Mandelbrot fractal.
package fractal

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"

	"github.com/andybons/gogif"
)

// MandelbrotParameters supplies parameters for the generation of a Mandelbrot
// image
type MandelbrotParameters struct {
	Xmin, Ymin, Xmax, Ymax, Width, Height, Contrast, Delay int
	Iterations, StartingIteration                          uint8
}

// WritePNG writes the Mandelbrot image to an io.Writer, encoded as a PNG
func WritePNG(w io.Writer, p MandelbrotParameters) {
	img := generateMandelbrot(p.Iterations, p)
	png.Encode(w, img)
}

// WriteAnimatedGIF writes an animation of the Mandelbrot fractal, increasing
// the number of iterations per frame
func WriteAnimatedGIF(w io.Writer, p MandelbrotParameters) {
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
	img := image.NewNRGBA(image.Rect(0, 0, p.Width, p.Height))
	for py := 0; py < p.Height; py++ {
		y := float64(py)/float64(p.Height)*float64(p.Ymax-p.Ymin) + float64(p.Ymin)
		for px := 0; px < p.Width; px++ {
			x := float64(px)/float64(p.Width)*float64(p.Xmax-p.Xmin) + float64(p.Xmin)
			z := complex(x, y)

			tries, escaped, zFinal := escapeIteration(z, iterations)
			shade := colorShade(tries, iterations, escaped, p.Contrast, zFinal)
			img.Set(px, py, shade)
		}
	}
	return img
}

func rgbaToPalleted(rgba image.Image) *image.Paletted {
	bounds := rgba.Bounds()
	palettedImage := image.NewPaletted(bounds, nil)
	quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	quantizer.Quantize(palettedImage, bounds, rgba, image.ZP)
	return palettedImage
}

func greyShade(tries, maxTries uint8, escaped bool, contrast int, zFinal complex128) color.Color {
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
	s := smooth(float64(tries), zFinal)
	hue := math.Abs((float64(s) / float64(maxTries)) * 360)
	log.Println("Hue : ", hue, "S: ", s, "z: ", zFinal)

	return HSVA{
		H: hue,
		S: 0.8,
		V: 1.0,
		A: 1.0,
	}
}

func smooth(iteration float64, zFinal complex128) float64 {
	return (iteration + 1) - (math.Log(math.Log(cmplx.Abs(zFinal))/math.Log(2)) / math.Log(2))
}

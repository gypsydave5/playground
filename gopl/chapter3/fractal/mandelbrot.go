// Package fractal presents functions to write fractal images to various image filetypes.
// It currently supports images in the Mandelbrot set as both PNG images and animated GIFs
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

var loggingEnabled bool

// MandelbrotParameters supplies parameters for the generation of a Mandelbrot
// image
type MandelbrotParameters struct {
	Xmin, Ymin, Xmax, Ymax, Width, Height, Contrast, Delay int
	Iterations, StartingIteration                          uint8
	Logging, Colour                                        bool
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

type shader func(tries, maxTries uint8, escaped bool, contrast int, zFinal complex128) color.Color

func animateMandelbrot(anim *gif.GIF, params MandelbrotParameters) {
	for i := params.StartingIteration; i < params.Iterations; i++ {
		img := generateMandelbrot(i, params)
		palettedImage := rgbaToPalleted(img)
		anim.Delay = append(anim.Delay, params.Delay)
		anim.Image = append(anim.Image, palettedImage)
	}
}

func generateMandelbrot(iterations uint8, params MandelbrotParameters) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, params.Width, params.Height))
	for py := 0; py < params.Height; py++ {
		for px := 0; px < params.Width; px++ {

			var shade color.Color
			if params.Colour == true {
				shade = superSample(iterations, px, py, params, colorShade)
			} else {
				shade = superSample(iterations, px, py, params, greyShade)
			}
			img.Set(px, py, shade)
		}
	}
	return img
}

func superSample(iterations uint8, px, py int, params MandelbrotParameters, sh shader) color.Color {
	const smoothing = 0.5
	x1, x2 := calcXs(px, smoothing, params)
	y1, y2 := calcYs(py, smoothing, params)

	zs := []complex128{complex(x1, y1), complex(x1, y2), complex(x2, y1), complex(x2, y2)}
	var colors []color.Color

	for _, z := range zs {
		t, e, zf := escapeIteration(z, iterations)
		colors = append(colors, sh(t, iterations, e, params.Contrast, zf))
	}

	return averageColor(colors...)
}

func calcXs(px int, smoothing float64, params MandelbrotParameters) (float64, float64) {
	adjustedWidth := float64(params.Width) * float64(params.Xmax-params.Xmin)
	return float64(px)/adjustedWidth + float64(params.Xmin), (float64(px)+smoothing)/adjustedWidth + float64(params.Xmin)
}

func calcYs(py int, smoothing float64, params MandelbrotParameters) (float64, float64) {
	adjustedLength := float64(params.Height) * float64(params.Ymax-params.Ymin)
	return float64(py)/adjustedLength + float64(params.Ymin), (float64(py)+smoothing)/adjustedLength + float64(params.Ymin)
}

func averageColor(colors ...color.Color) color.Color {
	length := float64(len(colors))
	var r, g, b, a float64
	for _, color := range colors {
		rc, gc, bc, ac := color.RGBA()
		r += float64(rc)
		g += float64(gc)
		b += float64(bc)
		a += float64(ac)
	}
	return color.RGBA64{
		uint16(r / length),
		uint16(g / length),
		uint16(b / length),
		uint16(a / length),
	}
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

	return smoothGrey(tries, maxTries, zFinal)
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

func smoothGrey(tries, maxTries uint8, zFinal complex128) color.Color {
	s := smooth(float64(tries), zFinal)
	shade := math.Abs((float64(s) / float64(maxTries)) * 255)
	return color.Gray{uint8(shade)}
}

func smoothHSV(tries, maxTries uint8, zFinal complex128) color.Color {
	s := smooth(float64(tries), zFinal)
	hue := math.Abs((float64(s) / float64(maxTries)) * 360)
	if loggingEnabled {
		log.Println("Hue : ", hue, "S: ", s, "z: ", zFinal)
	}

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

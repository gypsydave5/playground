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
				y1 := float64(py)/float64(params.Height)*float64(params.Ymax-params.Ymin) + float64(params.Ymin)
				y2 := (float64(py)+0.5)/float64(params.Height)*float64(params.Ymax-params.Ymin) + float64(params.Ymin)
				x1 := float64(px)/float64(params.Width)*float64(params.Xmax-params.Xmin) + float64(params.Xmin)
				x2 := (float64(px)+0.5)/float64(params.Width)*float64(params.Xmax-params.Xmin) + float64(params.Xmin)
				z1 := complex(x1, y1)
				z2 := complex(x1, y2)
				z3 := complex(x2, y1)
				z4 := complex(x2, y2)

				t1, e1, zf1 := escapeIteration(z1, iterations)
				t2, e2, zf2 := escapeIteration(z2, iterations)
				t3, e3, zf3 := escapeIteration(z3, iterations)
				t4, e4, zf4 := escapeIteration(z4, iterations)

				shade1 := colorShade(t1, iterations, e1, params.Contrast, zf1)
				shade2 := colorShade(t2, iterations, e2, params.Contrast, zf2)
				shade3 := colorShade(t3, iterations, e3, params.Contrast, zf3)
				shade4 := colorShade(t4, iterations, e4, params.Contrast, zf4)

				r1, g1, b1, a1 := shade1.RGBA()
				r2, g2, b2, a2 := shade2.RGBA()
				r3, g3, b3, a3 := shade3.RGBA()
				r4, g4, b4, a4 := shade4.RGBA()

				shade = color.RGBA64{
					uint16((r1 + r2 + r3 + r4) / 4),
					uint16((b1 + b2 + b3 + b4) / 4),
					uint16((g1 + g2 + g3 + g4) / 4),
					uint16((a1 + a2 + a3 + a4) / 4),
				}
			} else {
				y := float64(py)/float64(params.Height)*float64(params.Ymax-params.Ymin) + float64(params.Ymin)
				x := float64(px)/float64(params.Width)*float64(params.Xmax-params.Xmin) + float64(params.Xmin)
				z := complex(x, y)
				tries, escaped, zFinal := escapeIteration(z, iterations)
				shade = greyShade(tries, iterations, escaped, params.Contrast, zFinal)
			}
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

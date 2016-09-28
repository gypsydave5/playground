// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -4, -4, +4, +4
	width, height          = 1024, 1024
)

func main() {
	const iterations = 1
	const contrast = 15

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			// image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotShade(z, iterations, contrast))
		}
	}

	png.Encode(os.Stdout, img)
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

package fractal

import (
	"image/color"
	"math"
	"math/cmplx"
)

type shader func(tries, maxTries uint8, escaped bool, contrast int, zFinal complex128) color.Color

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

func smoothGrey(tries, maxTries uint8, zFinal complex128) color.Color {
	s := mandelbrotSmooth(float64(tries), zFinal)
	shade := math.Abs((float64(s) / float64(maxTries)) * 255)
	return color.Gray{uint8(shade)}
}

func smoothHSV(tries, maxTries uint8, zFinal complex128) color.Color {
	s := mandelbrotSmooth(float64(tries), zFinal)
	hue := math.Abs((float64(s) / float64(maxTries)) * 360)
	return HSVA{hue, 0.8, 1.0, 1.0}
}

func mandelbrotSmooth(iteration float64, zFinal complex128) float64 {
	return (iteration + 1) - (math.Log(math.Log(cmplx.Abs(zFinal))/math.Log(2)) / math.Log(2))
}

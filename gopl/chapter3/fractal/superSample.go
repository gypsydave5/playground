package fractal

import "image/color"

type pixelToColor func(vpixel) color.Color

func newMandelbrotPixelColorFunction(iterations uint8, params MandelbrotParameters, sh shader) pixelToColor {
	return func(p vpixel) color.Color {
		c := pixelToCoord(p, params)
		z := coordToComplex(c)
		return complexToMandelbrotColor(z, iterations, params, sh)
	}
}

func superSample(vp vpixel, pcFun pixelToColor) []color.Color {
	pixels := uniformGrid(vp)
	colors := make([]color.Color, len(pixels))
	for i, p := range pixels {
		colors[i] = pcFun(p)
	}
	return colors
}

func pixelColor(p vpixel, iterations uint8, params MandelbrotParameters, sh shader) color.Color {
	c := pixelToCoord(p, params)
	z := coordToComplex(c)
	return complexToMandelbrotColor(z, iterations, params, sh)
}

func zsToColors(zs []complex128, iterations uint8, params MandelbrotParameters, sh shader) []color.Color {
	colors := make([]color.Color, len(zs))
	for i, z := range zs {
		colors[i] = complexToMandelbrotColor(z, iterations, params, sh)
	}
	return colors
}

func complexToMandelbrotColor(z complex128, iterations uint8, params MandelbrotParameters, sh shader) color.Color {
	t, e, zf := escapeIteration(z, iterations)
	return sh(t, iterations, e, params.Contrast, zf)
}

func uniformGrid(vp vpixel) []vpixel {
	return []vpixel{
		vpixel{vp.X - 1, vp.Y - 1},
		vpixel{vp.X - 1, vp.Y + 1},
		vpixel{vp.X + 1, vp.Y - 1},
		vpixel{vp.X + 1, vp.Y + 1},
	}
}

func pixelsToCoords(ps []vpixel, params MandelbrotParameters) []coord {
	cs := make([]coord, len(ps))
	for i, p := range ps {
		cs[i] = pixelToCoord(p, params)
	}
	return cs
}

func coordsToComplexes(coords []coord) []complex128 {
	zs := make([]complex128, len(coords))
	for i, c := range coords {
		zs[i] = coordToComplex(c)
	}
	return zs
}

func coordToComplex(gc coord) complex128 {
	return complex(gc.X, gc.Y)
}

func pixelToCoord(vp vpixel, params MandelbrotParameters) coord {
	x := float64(vp.X)/float64(params.Width)*float64(params.Xmax-params.Xmin) + float64(params.Xmin)
	y := float64(vp.Y)/float64(params.Height)*float64(params.Ymax-params.Ymin) + float64(params.Ymin)
	return coord{x, y}
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

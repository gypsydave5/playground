package fractal

import (
	"image/color"
	"math/rand"
)

type pixelToColor func(vpixel) color.Color
type vpixel struct {
	X float64
	Y float64
}

func superSample(vp vpixel, pcFun pixelToColor) color.Color {
	pixels := uniformGridSample(vp)
	colors := make([]color.Color, len(pixels))
	for i, p := range pixels {
		colors[i] = pcFun(p)
	}
	return averageColor(colors...)
}

func uniformGridSample(vp vpixel) []vpixel {
	const d = 0.25
	return []vpixel{
		vpixel{vp.X - d, vp.Y - d},
		vpixel{vp.X - d, vp.Y + d},
		vpixel{vp.X + d, vp.Y - d},
		vpixel{vp.X + d, vp.Y + d},
	}
}

func randomSample(vp vpixel) []vpixel {
	const total = 8
	result := make([]vpixel, total)
	for i := range result {
		rx := (rand.Float64() * 2) - 1.0
		ry := (rand.Float64() * 2) - 1.0
		result[i] = vpixel{vp.X + rx, vp.Y + ry}
	}
	return result
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

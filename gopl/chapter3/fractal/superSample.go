package fractal

import "image/color"

type pixelToColor func(vpixel) color.Color
type vpixel struct {
	X float64
	Y float64
}

func superSample(vp vpixel, pcFun pixelToColor) color.Color {
	pixels := uniformGrid(vp)
	colors := make([]color.Color, len(pixels))
	for i, p := range pixels {
		colors[i] = pcFun(p)
	}
	return averageColor(colors...)
}

func uniformGrid(vp vpixel) []vpixel {
	const d = 0.25
	return []vpixel{
		vpixel{vp.X - d, vp.Y - d},
		vpixel{vp.X - d, vp.Y + d},
		vpixel{vp.X + d, vp.Y - d},
		vpixel{vp.X + d, vp.Y + d},
	}
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

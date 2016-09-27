package surface

import (
	"errors"
	"fmt"
	"image/color"
	"math"
)

var errInvalidColor = errors.New("Unable to parse hex color")

type rgbaByRange func(maxZ, minZ, z float64) color.RGBA

func newRGBAinRange(maxColorHex, minColorHex color.RGBA) rgbaByRange {
	hcFn := func(maxZ, minZ, z float64) color.RGBA {
		var ch color.RGBA
		ch.R = calculateMidByte(maxZ, minZ, maxColorHex.R, minColorHex.R, z)
		ch.G = calculateMidByte(maxZ, minZ, maxColorHex.G, minColorHex.G, z)
		ch.B = calculateMidByte(maxZ, minZ, maxColorHex.B, minColorHex.B, z)
		return ch
	}
	return hcFn
}

func calculateMidByte(maxFloat, minFloat float64, maxHex, minHex byte, f float64) byte {
	ratio := (f - minFloat) / (maxFloat - minFloat)
	hexRange := maxHex - minHex
	offset := minHex
	return byte(ratio*float64(hexRange) + float64(offset))
}

func rgbHexColorByRange(maxZ, minZ, z float64) color.RGBA {
	midPt := (maxZ + minZ) / 2
	var c color.RGBA

	r := math.Floor(255 * ((z - midPt) / (maxZ - midPt)))
	g := math.Floor(255 - math.Abs(255*(z/(maxZ-midPt))))
	b := math.Floor(255 * ((midPt - z) / (midPt - minZ)))

	if z > midPt {
		c.R = uint8(r)
	}
	if g > 0 {
		c.G = uint8(g)
	}
	if z < midPt {
		c.B = uint8(math.Abs(b))
	}

	return c
}

func rgbaToHex(c color.RGBA) string {
	return fmt.Sprintf("%02X%02X%02X", c.R, c.G, c.B)
}

package main

import (
	"fmt"
	"math"
)

type hexColorByRange func(maxZ, minZ, z float64) string

func newTestColorByRange(maxColor, minColor string) hexColorByRange {
	return func(maxZ, minZ, z float64) string {
		return "#FF0000"
	}
}

func rgbHexColorByRange(maxZ, minZ, z float64) string {
	midPt := (maxZ + minZ) / 2

	r := math.Floor(255 * ((z - midPt) / (maxZ - midPt)))
	g := math.Floor(255 - math.Abs(255*(z/(maxZ-midPt))))
	b := math.Floor(255 * ((midPt - z) / (midPt - minZ)))

	var bR, bG, bB byte

	if z > midPt {
		bR = byte(r)
	}
	if g > 0 {
		bG = byte(g)
	}
	if z < midPt {
		bB = byte(math.Abs(b))
	}

	color := fmt.Sprintf("#%02X%02X%02X", bR, bG, bB)
	return color
}

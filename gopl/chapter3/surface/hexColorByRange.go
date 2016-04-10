package main

import (
	"fmt"
	"math"
)

func hexColorByRange(maxZ, minZ, z float64) string {
	midPt := (maxZ - minZ) / 2
	r := math.Floor(255 * (z / midPt))
	g := math.Floor(255 - math.Abs(255*(z/midPt)))
	b := math.Floor(255 * (z / midPt))

	var bR, bG, bB byte
	if r > 0 {
		bR = byte(r)
	}
	bG = byte(g)
	if b < 0 {
		bB = byte(math.Abs(b))
	}

	color := fmt.Sprintf("#%02X%02X%02X", bR, bG, bB)
	return color
}

package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"
)

var errInvalidColor = errors.New("Unable to parse hex color")

type hexColorByRange func(maxZ, minZ, z float64) string
type colorHex struct {
	r byte
	g byte
	b byte
}

func (ch colorHex) String() string {
	b := []byte{ch.r, ch.g, ch.b}
	return strings.ToUpper(hex.EncodeToString(b))
}

func newTestColorByRange(maxColor, minColor string) (hexColorByRange, error) {
	var maxColorHex, minColorHex colorHex

	hcFn := func(maxZ, minZ, z float64) string {
		var ch colorHex
		ch.r = calculateMidByte(maxZ, minZ, maxColorHex.r, minColorHex.r, z)
		ch.g = calculateMidByte(maxZ, minZ, maxColorHex.g, minColorHex.g, z)
		ch.b = calculateMidByte(maxZ, minZ, maxColorHex.b, minColorHex.b, z)
		return ch.String()
	}

	maxColorHex, err := colorFromHexString(maxColor)
	if err != nil {
		return hcFn, err
	}
	minColorHex, err = colorFromHexString(minColor)
	if err != nil {
		return hcFn, err
	}

	return hcFn, nil
}

func calculateMidByte(maxFloat, minFloat float64, maxHex, minHex byte, f float64) byte {
	ratio := (f - minFloat) / (maxFloat - minFloat)
	hexRange := maxHex - minHex
	offset := minHex
	return byte(ratio*float64(hexRange) + float64(offset))
}

func colorFromHexString(s string) (colorHex, error) {
	var ch colorHex
	bytes, err := hex.DecodeString(s)
	if (err != nil) || (len(bytes) != 3) {
		return colorHex{}, errInvalidColor
	}
	ch.r = bytes[0]
	ch.g = bytes[1]
	ch.b = bytes[2]
	return ch, nil
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

	color := fmt.Sprintf("%02X%02X%02X", bR, bG, bB)
	return color
}

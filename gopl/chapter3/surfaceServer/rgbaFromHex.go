package main

import (
	"encoding/hex"
	"image/color"
)

func rgbaFromHex(colorHex string) (color.RGBA, error) {
	var c color.RGBA

	bytes, err := hex.DecodeString(colorHex)
	if err != nil {
		return c, err
	}

	c.R = bytes[0]
	c.G = bytes[1]
	c.B = bytes[2]
	return c, nil
}

package surface

import (
	"image/color"
	"testing"
)

// Tests around a custom colour function, calculating a colour in a range between
// two supplied colors.

var (
	red   = color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0)}
	blue  = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(255)}
	white = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255)}
	black = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0)}
)

func TestNewColorByRangeMax(t *testing.T) {
	maxColor := red
	minColor := blue

	hexColorByRange := newRGBAinRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := hexColorByRange(maxZ, minZ, z)

	expectedC := red
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestNewColorByRangeMin(t *testing.T) {
	maxColor := red
	minColor := blue

	hexColorByRange := newRGBAinRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := blue

	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestCalculateMidByte(t *testing.T) {
	maxFloat := 100.0
	minFloat := -100.0
	float := 0.0
	var maxHex byte = '\xFF'
	var minHex byte
	var expected byte = '\x7F'

	b := calculateMidByte(maxFloat, minFloat, maxHex, minHex, float)
	if b != expected {
		t.Errorf("Expected %X, but received %X", expected, b)
	}
}

func TestNewColorByRangeMid(t *testing.T) {
	maxColor := white
	minColor := black

	hexColorByRange := newRGBAinRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{127, 127, 127, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

// Tests surrounding the three color function, based on a midpoint of green,
// a max of red and a min of blue.
func TestZColorHexGreen(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{0, 255, 0, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRed(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{255, 0, 0, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRedAgain(t *testing.T) {
	maxZ := 0.9850673555377986
	minZ := -0.21285613860128652
	z := 0.9850673555377986
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{255, 0, 0, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexBlue(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{0, 0, 255, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.5
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{127, 127, 0, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexLowerMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -0.5
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := color.RGBA{0, 127, 127, 0}
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestRGBAToHex(t *testing.T) {
	var hex string

	hex = rgbaToHex(red)
	if hex != "FF0000" {
		t.Error("Expected 'FF0000', but got", hex)
	}

	hex = rgbaToHex(blue)
	if hex != "0000FF" {
		t.Error("Expected '0000FF', but got", hex)
	}

	hex = rgbaToHex(black)
	if hex != "000000" {
		t.Error("Expected '000000', but got", hex)
	}
}

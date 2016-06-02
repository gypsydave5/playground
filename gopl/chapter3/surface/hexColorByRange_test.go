package main

import "testing"

// Tests around a custom colour function, calculating a colour in a range between
// two supplied colors.
func TestNewColorByRangeMax(t *testing.T) {
	maxColor := "FF0000"
	minColor := "0000FF"

	hexColorByRange, _ := newTestColorByRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestNewColorByRangeMin(t *testing.T) {
	maxColor := "FF0000"
	minColor := "0000FF"

	hexColorByRange, _ := newTestColorByRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "0000FF"
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
	maxColor := "FFFFFF"
	minColor := "000000"

	hexColorByRange, _ := newTestColorByRange(maxColor, minColor)

	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := hexColorByRange(maxZ, minZ, z)
	expectedC := "7F7F7F"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestColorFromHexString(t *testing.T) {
	s := "FF0F10"
	c, _ := colorFromHexString(s)
	expectedColor := colorHex{
		r: '\xFF',
		g: '\x0F',
		b: '\x10',
	}
	if c != expectedColor {
		t.Errorf("Expected %v, yet we were given %v\n", expectedColor, c)
	}
}

func TestColorHexToString(t *testing.T) {
	c := colorHex{
		r: '\xFF',
		g: '\x0F',
		b: '\x10',
	}
	expected := "FF0F10"
	if c.String() != expected {
		t.Errorf("Expected %v, yet we were given %v\n", expected, c.String())
	}
}

// Tests surrounding the three color function, based on a midpoint of green,
// a max of red and a min of blue.
func TestZColorHexGreen(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "00FF00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRed(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRedAgain(t *testing.T) {
	maxZ := 0.9850673555377986
	minZ := -0.21285613860128652
	z := 0.9850673555377986
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexBlue(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "0000FF"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.5
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "7F7F00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexLowerMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -0.5
	c := rgbHexColorByRange(maxZ, minZ, z)
	expectedC := "007F7F"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

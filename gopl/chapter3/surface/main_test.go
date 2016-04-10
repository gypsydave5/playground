package main

import (
	"testing"
)
import "math"

func TestFMapperThrowsOnInfinity(t *testing.T) {
	_, _, _, err := fMapper(infinitude)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestFMapperThrowsOnNegInfinity(t *testing.T) {
	_, _, _, err := fMapper(negInfinitude)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestFMapperOKWithFinitude(t *testing.T) {
	_, _, _, err := fMapper(alwaysZero)(1, 1)
	if err != nil {
		t.Errorf("Shouldn't error on the finite")
	}
}

func TestZColorHexGreen(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.0
	c := zColorHex(maxZ, minZ, z)
	expectedC := "#00FF00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexRed(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 1.0
	c := zColorHex(maxZ, minZ, z)
	expectedC := "#FF0000"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexBlue(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := -1.0
	c := zColorHex(maxZ, minZ, z)
	expectedC := "#0000FF"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestZColorHexMid(t *testing.T) {
	maxZ := 1.0
	minZ := -1.0
	z := 0.5
	c := zColorHex(maxZ, minZ, z)
	expectedC := "#7F7F00"
	if c != expectedC {
		t.Errorf("Expected %v, got %v", expectedC, c)
	}
}

func TestNewPolygon(t *testing.T) {
	c := fMapper(alwaysZero)
	p := newPolygonGen(c)(0, 0)
	expectedP := polygon{
		ax: 302.5980762113533,
		ay: 11.5,
		bx: 300,
		by: 10,
		cx: 297.4019237886467,
		cy: 11.5,
		dx: 300,
		dy: 13,
	}
	if p != expectedP {
		t.Errorf("Expected p.z to be %#v, but got %#v", expectedP, p)
	}
}

func TestProject(t *testing.T) {
	sx, sy := project(0, 0, 0)
	if sx != 300 {
		t.Errorf("Expected sx to be 300, got %v", sx)
	}
	if sy != 160 {
		t.Errorf("Expected sx to be 160, got %v", sy)
	}
}

func infinitude(x, y float64) float64 {
	return math.Inf(1)
}

func negInfinitude(x, y float64) float64 {
	return math.Inf(-1)
}

func alwaysZero(x, y float64) float64 {
	return 0
}

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
		t.Errorf("Expected p to be %#v, but got %#v", expectedP, p)
	}
}

func TestGenerateSurface(t *testing.T) {
	c := fMapper(alwaysZero)
	p := newPolygonGen(c)
	surface := newSurface(p, 2)

	if len(surface.polygons) != 2 {
		t.Error("Expected polygon length to be 2, instead it's", len(surface.polygons))
	}
	if len(surface.polygons[0]) != 2 {
		t.Error("Expected polygon width to be 2, instead it's", len(surface.polygons[0]))
	}
	if surface.maxHeight != 0 {
		t.Error("Expected 0, got", surface.maxHeight)
	}
	if surface.minHeight != 0 {
		t.Error("Expected 0, got", surface.minHeight)
	}
}

func TestGenerateSVG(t *testing.T) {
	ps := make([][]polygon, 1)
	ps[0] = make([]polygon, 1)
	ps[0][0] = polygon{
		ax: 1.0,
		ay: 2.0,
		bx: 3.0,
		by: 4.0,
		cx: 5.0,
		cy: 6.0,
		dx: 7.0,
		dy: 8.0,
		z:  0,
	}
	s := surface{
		polygons:  ps,
		maxHeight: 1,
		minHeight: -1,
	}
	svg := generateSVG(s, 100, 200, rgbHexColorByRange).String()
	expectedSVGString := "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='100' height='200'>" +
		"<polygon points='1,2 3,4 5,6 7,8' fill='#00FF00'/>\n" +
		"</svg>"
	if svg != expectedSVGString {
		t.Errorf("Expected:\n\n%s\n\n, but got\n\n%s\n\n", expectedSVGString, svg)
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

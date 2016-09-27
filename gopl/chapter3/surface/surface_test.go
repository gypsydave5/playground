package surface

import (
	"bytes"
	"image/color"
	"math"
	"regexp"
	"testing"
)

func TestSurfaceFunctionMapperThrowsOnInfinity(t *testing.T) {
	_, _, _, err := surfaceFunctionMapper(infinitude, 3.0, 10)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestSurfaceFunctionMapperThrowsOnNegInfinity(t *testing.T) {
	_, _, _, err := surfaceFunctionMapper(negInfinitude, 3.0, 10)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestSurfaceFunctionMapperOKWithFinitude(t *testing.T) {
	_, _, _, err := surfaceFunctionMapper(alwaysZero, 3.0, 10)(1, 1)
	if err != nil {
		t.Errorf("Shouldn't error on the finite")
	}
}

func TestNewPolygon(t *testing.T) {
	c := surfaceFunctionMapper(alwaysZero, 30.0, 100)
	pf := newProjection(600, 320, 600/2.0/30.0, 320*0.4)
	p := polygonFactoryGenerator(c, pf)(0, 0)
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
	c := surfaceFunctionMapper(alwaysZero, 30.0, 100)
	pf := newProjection(600, 320, 320/2.0/30.0, 600*0.4)
	p := polygonFactoryGenerator(c, pf)
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
	svg := new(bytes.Buffer)
	generateSVG(s, 100, 200, rgbHexColorByRange, svg)
	expectedSVGString := "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='100' height='200'>" +
		"<polygon points='1,2 3,4 5,6 7,8' fill='#00FF00'/>\n" +
		"</svg>"
	if svg.String() != expectedSVGString {
		t.Errorf("Expected:\n\n%s\n\n, but got\n\n%s\n\n", expectedSVGString, svg.String())
	}
}

func TestNewOptions(t *testing.T) {
	red := color.RGBA{255, 0, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}

	opts := NewOptions()
	if opts.Cells != 100 {
		t.Error("Expected sensible default options to be generated")
	}
	if opts.Width != 600 {
		t.Error("Expected sensible default options to be generated")
	}
	if opts.Height != 320 {
		t.Error("Expected sensible default options to be generated")
	}
	if opts.XYRange != 30.0 {
		t.Error("Expected sensible default options to be generated")
	}
	if opts.UpperColor != red {
		t.Error("Expected sensible default options to be generated")
	}
	if opts.LowerColor != blue {
		t.Error("Expected sensible default options to be generated")
	}
}

func TestExportedSVG(t *testing.T) {
	b := new(bytes.Buffer)
	err := SVG(NewOptions(), b)
	if err != nil {
		t.Error("Expected SVG not to error")
	}
	regexpMatch, err := regexp.Match(`.*svg xmlns='http://www.w3.org/2000/svg'.*`, b.Bytes())
	if !regexpMatch {
		t.Errorf("Expected SVG to write an SVG string, but got \n\t%v", b)
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

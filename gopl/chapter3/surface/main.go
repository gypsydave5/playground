// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

type corner func(int, int) (float64, float64, float64, error)
type graphFun func(float64, float64) float64
type newPolygon func(int, int) polygon
type surface struct {
	polygons  [][]polygon
	maxHeight float64
	minHeight float64
}
type polygon struct {
	ax  float64
	ay  float64
	bx  float64
	by  float64
	cx  float64
	cy  float64
	dx  float64
	dy  float64
	z   float64
	err error
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	upperColor    = "00FF00"            // color of surface peaks (rgb hex)
	lowerColor    = "0000FF"            // color of surface troughs (rgb hex)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	corner := surfaceFunctionMapper(f)
	newPolygon := newPolygonGen(corner)
	surface := newSurface(newPolygon, cells)
	hexColorFunction, err := newTestColorByRange(upperColor, lowerColor)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, generateSVG(surface, width, height, hexColorFunction))
}

func newSurface(pFunc newPolygon, cells int) surface {
	var s surface
	s.minHeight = math.MaxFloat64
	s.maxHeight = -math.MaxFloat64
	s.polygons = make([][]polygon, cells)
	for i := 0; i < cells; i++ {
		s.polygons[i] = make([]polygon, cells)

		for j := 0; j < cells; j++ {
			p := pFunc(i, j)
			s.polygons[i][j] = p

			if p.z > s.maxHeight {
				s.maxHeight = p.z
			}
			if p.z < s.minHeight {
				s.minHeight = p.z
			}
		}
	}
	return s
}

func generateSVG(s surface, width int, height int, hcFn hexColorByRange) *bytes.Buffer {
	var b bytes.Buffer
	cells := len(s.polygons)
	b.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			b.WriteString(polygonToSVG(s.polygons[i][j], s.maxHeight, s.minHeight, hcFn))
		}
	}

	b.WriteString("</svg>")
	return &b
}

func polygonToSVG(p polygon, maxHeight float64, minHeight float64, hcFn hexColorByRange) string {
	color := hcFn(maxHeight, minHeight, p.z)
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%v'/>\n",
		p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, color)
}

func newPolygonGen(c corner) newPolygon {
	return func(i, j int) polygon {
		p := polygon{}
		x, y, z1, err := c(i+1, j)
		if err != nil {
			p.err = err
			return p
		}
		p.ax, p.ay = project(x, y, z1)
		x, y, z2, err := c(i, j)
		if err != nil {
			p.err = err
			return p
		}
		p.bx, p.by = project(x, y, z2)
		x, y, z3, err := c(i, j+1)
		if err != nil {
			p.err = err
			return p
		}
		p.cx, p.cy = project(x, y, z3)
		x, y, z4, err := c(i+1, j+1)
		if err != nil {
			p.err = err
			return p
		}
		p.dx, p.dy = project(x, y, z4)
		p.z = (z1 + z2 + z3 + z4) / 4
		return p
	}
}

func surfaceFunctionMapper(f graphFun) corner {
	return func(i, j int) (float64, float64, float64, error) {
		var err error
		// Find point (x,y) at corner of cell (i,j).
		x := xyrange * (float64(i)/cells - 0.5)
		y := xyrange * (float64(j)/cells - 0.5)

		// Compute surface height z.
		z := f(x, y)

		if math.Abs(z) == math.Inf(1) {
			err = fmt.Errorf("infinite result with args (%v, %v)", x, y)
			return x, y, z, err
		}

		// Project (x,y,z) isometrically onto a 2-D SVG canvas (sx,sy).
		return x, y, z, err
	}
}

func project(x, y, z float64) (sx, sy float64) {
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	//return (math.Sin(x) * math.Cos(y)) / 4
	return math.Cos(math.Abs(x)+math.Abs(y)) / 8
	//r := math.Hypot(x, y) // distance from (0,0)
	//return math.Sin(r) / r
}

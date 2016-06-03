// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

type corner func(int, int) (float64, float64, float64, error)
type graphFun func(float64, float64) float64
type surface struct {
	polygons  [][]polygon
	maxHeight float64
	minHeight float64
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
	project := newProjection(width, height, xyscale, zscale)
	polygonFactory := polygonFactoryGenerator(corner, project)
	surface := newSurface(polygonFactory, cells)
	maxColorHex, err := colorFromHexString(upperColor)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	minColorHex, err := colorFromHexString(lowerColor)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	hexColorFunction := newTestColorByRange(maxColorHex, minColorHex)
	io.Copy(os.Stdout, generateSVG(surface, width, height, hexColorFunction))
}

func newSurface(pFac polygonFactory, cells int) surface {
	var s surface
	s.minHeight = math.MaxFloat64
	s.maxHeight = -math.MaxFloat64
	s.polygons = make([][]polygon, cells)
	for i := 0; i < cells; i++ {
		s.polygons[i] = make([]polygon, cells)

		for j := 0; j < cells; j++ {
			p := pFac(i, j)
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

func f(x, y float64) float64 {
	//return (math.Sin(x) * math.Cos(y)) / 4
	return math.Cos(math.Abs(x)+math.Abs(y)) / 8
	//r := math.Hypot(x, y) // distance from (0,0)
	//return math.Sin(r) / r
}

// Package surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"image/color"
	"io"
	"math"
)

// GraphFunction takes an x and a y argument, and returns a z value. It describes
// a 3D surface
type GraphFunction func(x float64, y float64) (z float64)

// Options is the required arguments to generate a 3D surface from a function
//   function is the GraphFunction describing how to derive z from x and y values
//   cells is the number of cells that make up the displayed surface
//   width is the width of the canvas
//   height is the height of the canvas
//   xyrange is the range of x both and y, from -xyrange to +xyrange
//   upperColor is an RGB color represented as a hexadecimal string. It is the
// color of the cell with the highest calculated z value
//   lowerColor is an RGB color represented as a hexadecimal string. It is the
//  color of the cell with the lowest calculated z value
//
// All cells will be colored relative to their z value, calculating the appropriate
// intermediate color betewwn upperColor and lowerColor.
type Options struct {
	Function   GraphFunction
	Cells      int
	Width      int
	Height     int
	XYRange    float64
	UpperColor color.RGBA
	LowerColor color.RGBA
}
type corner func(int, int) (float64, float64, float64, error)
type surface struct {
	polygons  [][]polygon
	maxHeight float64
	minHeight float64
}

const angle = math.Pi / 6 // angle of x, y axes (=30°)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// SVG returns an io.Writer that writes an SVG of the 3D graph described by
// the function GraphFunction in the Options struct. See Options for more details.
func SVG(o Options, out io.Writer) error {
	xyscale := float64(o.Width) / 2.0 / o.XYRange // pixels per x or y unit
	zscale := float64(o.Height) * 0.4             // pixels per z unit

	corner := surfaceFunctionMapper(o.Function, o.XYRange, o.Cells)
	project := newProjection(o.Width, o.Height, xyscale, zscale)
	polygonFactory := polygonFactoryGenerator(corner, project)
	surface := newSurface(polygonFactory, o.Cells)

	maxColor := o.UpperColor
	minColor := o.LowerColor

	rgbaInRange := newRGBAinRange(maxColor, minColor)
	return generateSVG(surface, o.Width, o.Height, rgbaInRange, out)
}

// NewOptions returns an options struct for the surface with sane defaults set
func NewOptions() Options {
	red := color.RGBA{255, 0, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}

	return Options{
		Function:   defaultGraphFunction,
		Cells:      100,
		Width:      600,
		Height:     320,
		XYRange:    30.0,
		UpperColor: red,
		LowerColor: blue,
	}
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

func generateSVG(s surface, width int, height int, rgbaFn rgbaByRange, out io.Writer) error {
	cells := len(s.polygons)
	_, err := out.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)))
	if err != nil {
		return err
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			svgCell := []byte(
				polygonToSVG(s.polygons[i][j], s.maxHeight, s.minHeight, rgbaFn),
			)

			_, err = out.Write(svgCell)

			if err != nil {
				return err
			}
		}
	}

	_, err = out.Write([]byte("</svg>"))
	return err
}

func surfaceFunctionMapper(f GraphFunction, xyrange float64, cells int) corner {
	return func(i, j int) (float64, float64, float64, error) {
		var err error
		// Find point (x,y) at corner of cell (i,j).
		x := xyrange * (float64(i)/float64(cells) - 0.5)
		y := xyrange * (float64(j)/float64(cells) - 0.5)

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

func defaultGraphFunction(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

type corner func(int, int) (float64, float64, float64, error)
type graphFun func(float64, float64) float64

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	corner := fMapper(f)
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y, z, err := corner(i+1, j)
			if err != nil {
				continue
			}
			ax, ay := project(x, y, z)
			x, y, z, err = corner(i, j)
			if err != nil {
				continue
			}
			bx, by := project(x, y, z)
			x, y, z, err = corner(i, j+1)
			if err != nil {
				continue
			}
			cx, cy := project(x, y, z)
			x, y, z, err = corner(i+1, j+1)
			if err != nil {
				continue
			}
			dx, dy := project(x, y, z)

			colorR := math.Floor(255 * z)
			colorG := math.Floor(255 - math.Abs((255 * z)))
			colorB := math.Floor(255 * z)
			var bR, bG, bB byte

			if colorR > 0 {
				bR = byte(colorR)
			}
			if colorB < 0 {
				bB = byte(math.Abs(colorB))
			}
			bG = byte(colorG)

			color := fmt.Sprintf("#%02X%02x%02X", bR, bG, bB)
			fmt.Printf("<!-- (%v) %v %v %v -->", z, colorR, colorG, colorB)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%v'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func fMapper(f graphFun) corner {
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
	//return math.Cos(math.Abs(x)+math.Abs(y)) / 8
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

package surface

import "fmt"

type projectionFunction func(x, y, z float64) (sx, sy float64)
type polygonFactory func(int, int) polygon
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

func newProjection(width, height int, xyscale, zscale float64) projectionFunction {
	return func(x, y, z float64) (sx, sy float64) {
		sx = float64(width)/2.0 + (x-y)*cos30*xyscale
		sy = float64(height)/2.0 + (x+y)*sin30*xyscale - z*zscale
		return sx, sy
	}
}

func polygonToSVG(p polygon, maxHeight float64, minHeight float64, hcFn rgbaByRange) string {
	color := hcFn(maxHeight, minHeight, p.z)
	hex := rgbaToHex(color)
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%v'/>\n",
		p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, hex)
}

func polygonFactoryGenerator(c corner, project projectionFunction) polygonFactory {
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

package main

import "os"
import "github.com/gypsydave5/playground/gopl/chapter3/fractal"

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 8192, 8192
	maxIterations          = 40
	startingIteration      = 39
	contrast               = 15
	delay                  = 20
)

func main() {
	params := fractal.MandelbrotParameters{
		Xmin:       -2,
		Ymin:       -2,
		Xmax:       2,
		Ymax:       2,
		Width:      256,
		Height:     256,
		Iterations: 20,
		Contrast:   15,
		Delay:      20,
	}
	fractal.WritePNG(os.Stdout, params)
}

package main

import (
	"flag"
	"os"

	"github.com/gypsydave5/playground/gopl/chapter3/fractal"
)

var (
	bounds                        float64
	width, height                 int
	iterations, startingIteration uint
	format                        string
	logging, colour, supersample  bool
)

func init() {
	flag.Float64Var(&bounds, "bounds", 2, "max and min for x and y axes")
	flag.IntVar(&width, "width", 256, "image width")
	flag.IntVar(&height, "height", 256, "image height")
	flag.UintVar(&iterations, "iterations", 40, "max iterations to perform to see if point escapes")
	flag.UintVar(&startingIteration, "startingIteration", 0, "animated only - iterations to start from")
	flag.StringVar(&format, "format", "png", "output format - defaults to png. Set to 'gif' for animated gif")
	flag.BoolVar(&logging, "verbose", false, "output log messages to stderr")
	flag.BoolVar(&colour, "colour", true, "output in colour or greyscale.")
	flag.BoolVar(&supersample, "ss", false, "turns supersampling per pixel on or off")
}

func main() {
	flag.Parse()

	params := fractal.MandelbrotParameters{
		Xmin:              -bounds,
		Ymin:              -bounds,
		Xmax:              bounds,
		Ymax:              bounds,
		Width:             width,
		Height:            height,
		Iterations:        uint8(iterations),
		StartingIteration: uint8(startingIteration),
		Contrast:          15,
		Delay:             20,
		Logging:           logging,
		Colour:            colour,
		SuperSample:       supersample,
	}

	fractal.WritePNG(os.Stdout, params)
}

package main

import (
	"flag"
	"log"
	"math"
	"os"

	"github.com/gypsydave5/playground/gopl/chapter3/surface"
)

var width = flag.Int("width", 600, "width of the SVG canvas")
var height = flag.Int("height", 320, "height of the SVG canvas")
var cells = flag.Int("cells", 100, "number of cells in the SVG")
var xyrange = flag.Float64("range", 30.0, "range of x and y axes (-range..+range)")
var upperColor = flag.String("upperColor", "FF0000", "color of maximum z value polygon")
var lowerColor = flag.String("lowerColor", "0000FF", "color of minimum z value polygon")
var outputFile = flag.String("out", "", "name of file to output SVG to")

func main() {
	flag.Parse()

	if *outputFile == "" {
		err := surface.SVG(functionOne, *cells, *width, *height, *xyrange, *upperColor, *lowerColor, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	err = surface.SVG(functionOne, *cells, *width, *height, *xyrange, *upperColor, *lowerColor, file)
	if err != nil {
		log.Fatal(err)
	}

}

func functionOne(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func functionTwo(x, y float64) float64 {
	return math.Cos(math.Abs(x)+math.Abs(y)) / 8
}

func functionThree(x, y float64) float64 {
	return (math.Sin(x) * math.Cos(y)) / 4
}

package main

import (
	"flag"
	"log"
	"math/rand"
	"os"

	"github.com/gypsydave5/playground/gopl/chapter1/lissajous"
)

var cycles = flag.Float64("cycles", 5.0, "number of complete x oscillator revolutions")
var resolution = flag.Float64("resolution", 0.001, "angular resolution")
var size = flag.Int("size", 100, "image canvas covers [-size..+size]")
var frames = flag.Int("frames", 64, "number of animation frames")
var delay = flag.Int("delay", 8, "delay between frames in 10ms units")
var frequency = flag.Float64("frequency", rand.Float64()*3, "relative frequency of y oscillator")
var outputFile = flag.String("out", "", "name of file to output GIF to")

func main() {
	flag.Parse()
	opts := lissajous.Options{
		Cycles:     *cycles,
		Resolution: *resolution,
		Size:       *size,
		Frames:     *frames,
		Delay:      *delay,
		Frequency:  *frequency,
	}

	if *outputFile == "" {
		lissajous.Lissajous(os.Stdout, opts)
	} else {
		f, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		} else {
			lissajous.Lissajous(f, opts)
		}
	}
}

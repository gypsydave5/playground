package main

import (
	"github.com/gypsydave5/playground/gopl/lissajous"
	"os"
)

func main() {
	opts := lissajous.Options{
		Cycles:     5,
		Resolution: 0.001,
		Size:       100,
		Frames:     64,
		Delay:      8,
	}
	lissajous.Lissajous(os.Stdout, opts)
}

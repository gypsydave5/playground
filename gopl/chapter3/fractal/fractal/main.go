package main

import "os"
import "github.com/gypsydave5/playground/gopl/chapter3/fractal"

func main() {
	fractal.WritePNG(os.Stdout)
}

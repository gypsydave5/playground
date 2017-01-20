package main

import (
	"log"
	"net/http"

	"github.com/gypsydave5/playground/gopl/chapter3/fractal"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	var bounds = 2
	var width, height = 1024, 1024
	var colour = true
	var supersample = true
	var logging = false
	var iterations uint8 = 40

	opts := fractal.MandelbrotParameters{
		Xmin:        -bounds,
		Ymin:        -bounds,
		Xmax:        bounds,
		Ymax:        bounds,
		Width:       width,
		Height:      height,
		Iterations:  iterations,
		Contrast:    15,
		Delay:       20,
		Logging:     logging,
		Colour:      colour,
		SuperSample: supersample,
	}
	log.Printf("Options: %v", opts)
	fractal.WritePNG(w, opts)
}

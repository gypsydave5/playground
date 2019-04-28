package main

import (
	"log"
	"net/http"

	"github.com/gypsydave5/playground/gopl/chapter3/fractal"
)

func main() {
	http.HandleFunc("/", router)
	log.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL: %+v\n", r.URL.Path)

	switch r.URL.Path {
	case "/mandlebrot":
		fractalHandler(w, r)
	default:
		viewerHandler(w, r)
	}
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var x, y float64 = 0, 0
	var width, height = 1024, 1024
	var colour = true
	var supersample = true
	var logging = false
	var iterations uint8 = 40

	coord := fractal.Coord{X: x, Y: y}
	bounds := *fractal.CoordsZoomToBounds(coord, 1, 2)

	opts := fractal.Parameters{
		Bounds:      bounds,
		Width:       width,
		Height:      height,
		Iterations:  iterations,
		Contrast:    15,
		Delay:       20,
		Logging:     logging,
		Colour:      colour,
		SuperSample: supersample,
	}

	opts = applyOptions(opts, r.Form)

	log.Printf("Options: %+v", opts)
	fractal.WritePNG(w, opts)
}

func viewerHandler(w http.ResponseWriter, r *http.Request) {
	t := viewer()
	t.Execute(w, nil)
}

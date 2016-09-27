package main

import (
	"log"
	"net/http"

	"github.com/gypsydave5/playground/gopl/chapter3/surface"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	opts := surface.NewOptions()

	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	surface.SVG(opts, w)
}

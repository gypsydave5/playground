package main

import (
	"github.com/gypsydave5/playground/gopl/lissajous"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	opts := lissajous.Options{}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	opts = lissajousOpts(r.Form)
	log.Printf("Options: %v", opts)
	lissajous.Lissajous(w, opts)
}

func lissajousOpts(f url.Values) lissajous.Options {
	opts := lissajous.Options{}
	setDefaultFloat64(&opts.Cycles, f["cycles"], 5.0)
	setDefaultFloat64(&opts.Resolution, f["resolution"], 0.001)
	setDefaultFloat64(&opts.Frequency, f["frequency"], rand.Float64()*3)
	setDefaultInt(&opts.Delay, f["delay"], 8)
	setDefaultInt(&opts.Frames, f["frames"], 64)
	setDefaultInt(&opts.Size, f["size"], 100)
	return opts
}

func setDefaultInt(o *int, v []string, d int) {
	if v != nil {
		*o, _ = strconv.Atoi(v[0])
	} else {
		*o = d
	}
}

func setDefaultFloat64(o *float64, v []string, d float64) {
	if v != nil {
		*o, _ = strconv.ParseFloat(v[0], 64)
	} else {
		*o = d
	}
}

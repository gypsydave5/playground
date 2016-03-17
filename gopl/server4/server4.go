package main

import (
	"github.com/gypsydave5/playground/gopl/lissajous"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	opts := lissajous.Options{
		Cycles:     5,
		Resolution: 0.001,
		Size:       100,
		Frames:     64,
		Delay:      8,
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	opts = lissajousOpts(r.Form)
	lissajous.Lissajous(w, opts)
}

func lissajousOpts(f url.Values) lissajous.Options {
	opts := lissajous.Options{}
	setDefaultFloat64(&opts.Cycles, f, "cycles", 5.0)
	setDefaultFloat64(&opts.Resolution, f, "resolution", 0.001)
	setDefaultInt(&opts.Size, f, "size", 100)
	return opts
}

func setDefaultInt(o *int, f url.Values, k string, d int) {
	if f[k] != nil {
		*o, _ = strconv.Atoi(f[k][0])
	} else {
		*o = d
	}
}

func setDefaultFloat64(o *float64, f url.Values, k string, d float64) {
	if f[k] != nil {
		*o, _ = strconv.ParseFloat(f[k][0], 64)
	} else {
		*o = d
	}
}

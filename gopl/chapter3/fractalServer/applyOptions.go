package main

import (
	"log"
	"net/url"
	"strconv"

	"github.com/gypsydave5/playground/gopl/chapter3/surface"
)

func applyOptions(opts surface.Options, params url.Values) surface.Options {
	if _, pres := params["cells"]; pres {
		cells, err := strconv.Atoi(params["cells"][0])
		if err != nil {
			log.Println("unexpected parameter [cells]: ", params["cells"][0])
		} else {
			opts.Cells = cells
		}
	}
	return opts
}

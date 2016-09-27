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

	if _, pres := params["width"]; pres {
		width, err := strconv.Atoi(params["width"][0])
		if err != nil {
			log.Println("unexpected parameter [width]: ", params["width"][0])
		} else {
			opts.Width = width
		}
	}

	if _, pres := params["height"]; pres {
		height, err := strconv.Atoi(params["height"][0])
		if err != nil {
			log.Println("unexpected parameter [height]: ", params["height"][0])
		} else {
			opts.Height = height
		}
	}

	if _, pres := params["xyrange"]; pres {
		xyrange, err := strconv.ParseFloat(params["xyrange"][0], 32)
		if err != nil {
			log.Println("unexpected parameter [xyrange]: ", params["xyrange"][0])
		} else {
			opts.XYRange = xyrange
		}
	}

	if _, pres := params["lowercolor"]; pres {
		lowercolor := params["lowercolor"][0]
		rgba, err := rgbaFromHex(lowercolor)
		if err != nil {
			log.Println("unexpected parameter [lowercolor]: ", params["lowercolor"][0])
		} else {
			opts.LowerColor = rgba
		}
	}

	if _, pres := params["uppercolor"]; pres {
		uppercolor := params["uppercolor"][0]
		rgba, err := rgbaFromHex(uppercolor)
		if err != nil {
			log.Println("unexpected parameter [uppercolor]: ", params["uppercolor"][0])
		} else {
			opts.UpperColor = rgba
		}
	}

	return opts
}

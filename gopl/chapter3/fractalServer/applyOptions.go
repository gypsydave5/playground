package main

import (
	"log"
	"net/url"
	"strconv"

	"github.com/gypsydave5/playground/gopl/chapter3/fractal"
)

func applyOptions(opts fractal.Parameters, params url.Values) fractal.Parameters {
	zoom := 1.0
	defaultBounds := 2.0

	var coord = fractal.Coord{}

	if _, pres := params["x"]; pres {
		x, err := strconv.ParseFloat(params["x"][0], 64)
		if err != nil {
			log.Println("unexpected parameter [x]: ", params["x"][0])
		}
		coord.X = x
	}

	if _, pres := params["y"]; pres {
		y, err := strconv.ParseFloat(params["y"][0], 64)
		if err != nil {
			log.Println("unexpected parameter [y]: ", params["y"][0])
		}
		coord.Y = y
	}

	if _, pres := params["zoom"]; pres {
		z, err := strconv.ParseFloat(params["zoom"][0], 64)
		if err != nil {
			log.Println("unexpected parameter [zoom]: ", params["zoom"][0])
		}
		zoom = z
	}

	if _, pres := params["iterations"]; pres {
		iterations, err := strconv.Atoi(params["iterations"][0])
		if err != nil {
			log.Println("unexpected parameter [iterations]: ", params["iterations"][0])
		}
		opts.Iterations = uint8(iterations)
	}

	if _, pres := params["width"]; pres {
		width, err := strconv.Atoi(params["width"][0])
		if err != nil {
			log.Println("unexpected parameter [width]: ", params["width"][0])
		}
		opts.Width = width
	}

	if _, pres := params["height"]; pres {
		height, err := strconv.Atoi(params["height"][0])
		if err != nil {
			log.Println("unexpected parameter [height]: ", params["height"][0])
		}
		opts.Height = height
	}

	if _, pres := params["colour"]; pres {
		colour, err := strconv.ParseBool(params["colour"][0])
		if err != nil {
			log.Println("unexpected parameter [colour]: ", params["colour"][0])
		}
		opts.Colour = colour
	}

	if _, pres := params["supersample"]; pres {
		supersample, err := strconv.ParseBool(params["supersample"][0])
		if err != nil {
			log.Println("unexpected parameter [supersample]: ", params["supersample"][0])
		}
		opts.SuperSample = supersample
	}

	bounds := fractal.CoordsZoomToBounds(coord, zoom, defaultBounds)

	opts.Bounds = *bounds

	return opts
}

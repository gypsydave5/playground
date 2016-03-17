package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
)

//Options provides options for the Lissajous generator
type Options struct {
	Cycles     float64
	Resolution float64
	Size       int
	Frames     int
	Delay      int
}

//Lissajous generates a Lissajous animation
func Lissajous(out io.Writer, o Options) {
	cycles := o.Cycles           // number of complete x oscillator revolutions
	res := o.Resolution          // angular resolution
	size := o.Size               // image canvas covers [-size..+size]
	nframes := o.Frames          // number of animation frames
	delay := o.Delay             // delay between frames in 10ms units
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		var colorIndex uint8
		if i%2 == 0 {
			colorIndex = greenIndex
		} else {
			colorIndex = redIndex
		}
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

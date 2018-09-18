package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{128, 0, 0, 128},
	color.RGBA{0, 128, 0, 128},
	color.RGBA{0, 0, 128, 128},
}

const whiteIndex = 0 // first color in palette
const blackIndex = 1 // second color in palette

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if cycle := r.FormValue("cycles"); cycle != "" {
		cycles, _ = strconv.Atoi(cycle)
	}

	size := 100
	if cover := r.FormValue("size"); cover != "" {
		size, _ = strconv.Atoi(cover)
	}

	lissajous(w, cycles, size)
}

func lissajous(out io.Writer, cycles, size int) {
	const (
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase differences

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			c := uint8(rand.Intn(3) + 1)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), c)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

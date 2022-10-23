package main

import (
	"fmt"
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

var palette = []color.Color{color.RGBA{122, 182, 205, 1}, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type Parameters struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {

		p := Parameters{cycles: 5, res: 0.0001, size: 100, nframes: 64, delay: 8}

		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		for k, v := range r.Form {
			if k == "cycle" {
				p.cycles, _ = strconv.Atoi(v[0])
			} else if k == "res" {
				p.res, _ = strconv.ParseFloat(v[0], 64)
			} else if k == "size" {
				p.size, _ = strconv.Atoi(v[0])
			} else if k == "nframes" {
				p.nframes, _ = strconv.Atoi(v[0])
			} else if k == "delay" {
				p.delay, _ = strconv.Atoi(v[0])
			}
		}

		lisssajous(w, p)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func lisssajous(out io.Writer, p Parameters) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*float64(math.Pi); t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding error
}

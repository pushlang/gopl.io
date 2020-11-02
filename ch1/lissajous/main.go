// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
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
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var colorBack = color.RGBA{0x02, 0x14, 0x01, 0xFF}
var colorFore = color.RGBA{0x36, 0xb8, 0x33, 0xFF}

var palette = []color.Color{colorBack, colorFore}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, nil)
}

func lissajous(out io.Writer, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles, nframes, size, delay := 5., 64, 100., 8
	//?cycles=5&nframes=64&size=100&delay=8

	for k, v := range r.Form {
		switch k {
		case "cycles":
			a, _ := strconv.Atoi(v[0])
			cycles = float64(a)
		case "size":
			a, _ := strconv.Atoi(v[0])
			size = float64(a)
		case "nframes":
			nframes, _ = strconv.Atoi(v[0])
		case "delay":
			delay, _ = strconv.Atoi(v[0])
		}
		//fmt.Fprintf(out, "Form[%q] = %q\n", k, v)
	}

	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		//size = 100   // image canvas covers [-size..+size]
		//nframes = 64    // number of animation frames
		//delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	cb := color.RGBA{uint8(rand.Int() * 255), uint8(rand.Int() * 255), uint8(rand.Int() * 255), 0xFF}
	cf := color.RGBA{uint8(rand.Int() * 255), uint8(rand.Int() * 255), uint8(rand.Int() * 255), 0xFF}
	palette = []color.Color{cb, cf}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))

		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	crand "crypto/rand"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	rand "math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func randBytes(n int) []uint8 {
	b := make([]byte, n)
	crand.Read(b)

	return []uint8(b)
}

func randColor() (b color.Color, f color.Color) {
	RGBb := randBytes(3)
	//RGBf := make([]uint8, 3)

	// for i, v := range RGBb {
	// 	for j := 0; j < 8; j++ {
	// 		RGBf[i] = uint8(v) ^ (1 << j)
	// 	}
	// }

	b = color.RGBA{RGBb[0], RGBb[1], RGBb[2], 0xff}
	f = color.RGBA{^RGBb[0], ^RGBb[1], ^RGBb[2], 0xff}
	//f = color.RGBA{RGBf[0], RGBf[1], RGBf[2], 0xff}

	return
}

func parseParametrsURL(u string) map[string][]string {
	pu, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(pu.RawQuery)

	return m
}

//!+main
var (
	backgroundColor, foregroundColor color.Color
)

// var palette = []color.Color{color.White, color.Black}
var palette []color.Color

const (
	backgroundIndex = 0 // first color in palette
	foregroundIndex = 1 // next color in palette
)

func main() {
	backgroundColor, foregroundColor = randColor()
	palette = []color.Color{backgroundColor, foregroundColor}

	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			c := parseParametrsURL(r.URL.RequestURI())
			//fmt.Fprintf(w, "cycles = %d\n", c)
			lissajous(w, c)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main ?cycles=1&res=1&size=100&nframes=128&delay=4
	m := map[string][]string{"cycles": {"3"}, "res": {"1"}, "size": {"200"}, "nframes": {"128"}, "delay": {"4"}}
	lissajous(os.Stdout, m)
}

func lissajous(out io.Writer, p map[string][]string) {

	var (
		cycles, _  = strconv.Atoi(p["cycles"][0])  // number of complete x oscillator revolutions
		ress, _    = strconv.Atoi(p["res"][0])     //0.001 // angular resolution
		sizee, _   = strconv.Atoi(p["size"][0])    //200   // image canvas covers [-size..+size]
		nframes, _ = strconv.Atoi(p["nframes"][0]) //128   // number of animation frames
		delay, _   = strconv.Atoi(p["delay"][0])   //4     // delay between frames in 10ms units
	)
	res := float64(ress) / 1000.
	size := float64(sizee)

	freq := rand.Float64() * 3 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	_, foregroundColor = randColor()
	palette = []color.Color{backgroundColor, foregroundColor}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5),
				foregroundIndex)
		}
		phase += 0.01
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main

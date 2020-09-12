// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			createSVG(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	//createSVG(os.Stdout)
}

func createSVG(out io.Writer) {
	fmt.Fprintf(out,
		//"<?xml version=\"1.0\"?>"+
		"<svg xmlns=\"http://www.w3.org/2000/svg\"\n"+
			//"xmlns:xlink=\"http://www.w3.org/1999/xlink\">\n"+
			"style=\"stroke: grey; fill: white; stroke-width: 0.7\" \n"+
			"width=\"%d\" height=\"%d\">\n", width, height)

	drawPolygons(out)

	fmt.Fprintf(out, "</svg>")
}

func drawPolygons(out io.Writer) {
	zDelta, _, zMin := getMinMaxZ()
	//fmt.Printf("Max:%f\tMin:%f\tDelta:%f", zMax, zMin, zDelta)
	color := ""
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			color = chooseColor(color, dz, zDelta, zMin)

			fmt.Fprintf(out, "<polygon points=\"%g,%g %g,%g %g,%g %g,%g\" stroke=\"%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
}

func chooseColor(color string, dz float64, zDelta float64, zMin float64) string {
	switch {
	case dz < (zMin + zDelta/5):
		color = "#0000ff"
	case (zMin+zDelta/5) < dz && dz < (zMin+zDelta/4):
		color = "#00cc00"
	case (zMin+zDelta/4) < dz && dz < (zMin+zDelta/3):
		color = "#ffff00"
	case (zMin+zDelta/3) < dz && dz < (zMin+zDelta/2):
		color = "#ff6600"
	case (zMin+zDelta/2) < dz && dz < (zMin+zDelta/1):
		color = "#ff0000"
	}
	return color
}

func getMinMaxZ() (float64, float64, float64) {
	zMax := 0.
	zMin := 0.
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := corner(i, j)
			if z > zMax {
				zMax = z
			}
			if z < zMin {
				zMin = z
			}
		}
	}
	zDelta := zMax - zMin
	return zDelta, zMax, zMin
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y, float64(i), float64(j))
	if math.IsNaN(z) || math.IsInf(z, 1) || math.IsInf(z, -1) {
		z = 0
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y, i, j float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return (math.Sin(r * 1.0)) / r
}
//!-

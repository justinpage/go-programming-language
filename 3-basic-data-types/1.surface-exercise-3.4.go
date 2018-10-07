// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            //  canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixes per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var sideToSide, topToBottom int

type Points struct {
	X, Y    float64
	Valleys bool
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sideToSide = width
	if s := r.FormValue("width"); s != "" {
		sideToSide, _ = strconv.Atoi(s)
	}

	topToBottom = 320
	if h := r.FormValue("height"); h != "" {
		topToBottom, _ = strconv.Atoi(h)
	}

	color := "white"
	if c := r.FormValue("color"); c != "" {
		color = c
	}

	surface(w, sideToSide, topToBottom, color)
}

func surface(w http.ResponseWriter, sideToSide int, topToBottom int, color string) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			pointsA, err := corner(i+1, j)
			if err != nil {
				continue
			}

			pointsB, err := corner(i, j)
			if err != nil {
				continue
			}

			pointsC, err := corner(i, j+1)
			if err != nil {
				continue
			}

			pointsD, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				pointsA.X, pointsA.Y, pointsB.X, pointsB.Y,
				pointsC.X, pointsC.Y, pointsD.X, pointsD.Y)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (*Points, error) {
	// Find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height of z
	z := f(x, y)

	if math.IsNaN(z) {
		return &Points{}, errors.New("Say no to infiniti")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(sideToSide)/2 + (x-y)*cos30*xyscale
	sy := float64(topToBottom)/2 + (x+y)*sin30*xyscale - z*zscale

	return &Points{sx, sy, false}, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

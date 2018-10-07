// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"math"
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

type Points struct {
	X, Y    float64
	Valleys bool
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
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

			if pointsA.Valleys {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
					"style='fill:blue'/>\n",
					pointsA.X, pointsA.Y, pointsB.X, pointsB.Y,
					pointsC.X, pointsC.Y, pointsD.X, pointsD.Y)
			} else {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
					"style='fill:red'/>\n",
					pointsA.X, pointsA.Y, pointsB.X, pointsB.Y,
					pointsC.X, pointsC.Y, pointsD.X, pointsD.Y)
			}
		}
	}
	fmt.Println("</svg>")
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
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if z < 0 {
		return &Points{sx, sy, true}, nil
	}

	return &Points{sx, sy, false}, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

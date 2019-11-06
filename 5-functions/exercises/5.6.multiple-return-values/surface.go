// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes  (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type point struct {
	x, y float64
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			pointA, err := corner(i+1, j)
			if err != nil {
				continue
			}

			pointB, err := corner(i, j)
			if err != nil {
				continue
			}

			pointC, err := corner(i, j+1)
			if err != nil {
				continue
			}

			pointD, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Printf(
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				pointA.x, pointA.y, pointB.x, pointB.y,
				pointC.x, pointC.y, pointD.x, pointD.y,
			)
		}
	}

	fmt.Printf("</svg>")
}

func corner(i, j int) (p *point, err error) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsNaN(z) {
		err = errors.New("not a number")
		return
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	p = &point{sx, sy}
	return
}

func f(x, y float64) float64 {
	r := math.Cos(x) // distance from (0,0)
	return (math.Sin(r) / r) * math.Cos(y)
}

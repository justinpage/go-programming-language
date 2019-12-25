package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/justinpage/go-programming-language/7-interfaces/examples/9.expression-evaluator/eval"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes  (=30°)
)

type point struct {
	x, y float64
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func surface(out http.ResponseWriter, f func(x, y float64) float64) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			pointA, err := corner(i+1, j, f)
			if err != nil {
				continue
			}

			pointB, err := corner(i, j, f)
			if err != nil {
				continue
			}

			pointC, err := corner(i, j+1, f)
			if err != nil {
				continue
			}

			pointD, err := corner(i+1, j+1, f)
			if err != nil {
				continue
			}

			fmt.Fprintf(out,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				pointA.x, pointA.y, pointB.x, pointB.y,
				pointC.x, pointC.y, pointD.x, pointD.y,
			)
		}
	}

	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (*point, error) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsNaN(z) {
		return &point{}, errors.New("not a number")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale

	if z > 0 {
		return &point{sx, sy}, nil
	}

	return &point{sx, sy}, nil
}

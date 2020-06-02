// Surface computes an SVG that renders a 3-D surface within the browser.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

const (
	width, height = 600, 320    // canvas size in pixels
	cells         = 100         // number of grid cells
	xyrange       = 30.0        // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6 // angle of x, y axes  (=30°)
)

type options struct {
	width   int
	height  int
	color   string
	xyscale float64
	zscale  float64
}

type point struct {
	key  string
	x, y float64
}

type square struct {
	pointA, pointB, pointC, pointD *point
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		option := &options{
			width, height, "white",
			(width / 2 / xyrange), (height * 0.4),
		}

		if w := r.Form.Get("width"); w != "" {
			option.width, _ = strconv.Atoi(w)
			option.xyscale = (float64(option.width) / 2 / xyrange)
		}

		if h := r.Form.Get("height"); h != "" {
			option.height, _ = strconv.Atoi(h)
			option.zscale = (float64(option.height) * 0.4)
		}

		if c := r.Form.Get("color"); c != "" {
			option.color = c
		}

		surface(w, option)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func surface(out http.ResponseWriter, opt *options) {
	out.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", opt.color, opt.width, opt.height)

	var wg1 sync.WaitGroup
	squares := make(chan *square)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var wg2 sync.WaitGroup
			ch := make(chan *point)

			wg2.Add(4)
			go corner("A", i+1, j, opt, ch, &wg2)
			go corner("B", i, j, opt, ch, &wg2)
			go corner("C", i, j+1, opt, ch, &wg2)
			go corner("D", i+1, j+1, opt, ch, &wg2)

			go func() {
				wg2.Wait()
				close(ch)
			}()

			wg1.Add(1)
			go func() {
				defer wg1.Done()
				sq := &square{}
				for p := range ch {
					switch p.key {
					case "A":
						sq.pointA = p
					case "B":
						sq.pointB = p
					case "C":
						sq.pointC = p
					case "D":
						sq.pointD = p
					default:
						return
					}
				}

				squares <- sq
			}()
		}
	}

	go func() {
		wg1.Wait()
		close(squares)
	}()

	for sq := range squares {
		fmt.Fprintf(out,
			"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			sq.pointA.x, sq.pointA.y, sq.pointB.x, sq.pointB.y,
			sq.pointC.x, sq.pointC.y, sq.pointD.x, sq.pointD.y,
		)
	}

	fmt.Fprintf(out, "</svg>")
}

func corner(key string, i, j int, opt *options,
	ch chan *point, wg *sync.WaitGroup,
) {
	defer wg.Done()
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	// Ignore for now
	if math.IsNaN(z) {
		// 	return &point{}, errors.New("not a number")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := float64(opt.width/2) + (x-y)*cos30*opt.xyscale
	sy := float64(opt.height/2) + (x+y)*sin30*opt.xyscale - z*opt.zscale

	if z > 0 {
		ch <- &point{key, sx, sy}
		return
	}

	ch <- &point{key, sx, sy}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

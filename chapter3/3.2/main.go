package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var function = saddle

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, infinite := corner(i+1, j, function)
			if infinite {
				continue
			}

			bx, by, infinite := corner(i, j, function)
			if infinite {
				continue
			}

			cx, cy, infinite := corner(i, j+1, function)
			if infinite {
				continue
			}

			dx, dy, infinite := corner(i+1, j+1, function)
			if infinite {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, function func(float64, float64) float64) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := function(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0.0, 0.0, true
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, false
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	const a = 0.1
	const b = 1.0
	return a * (math.Sin(x/b) + math.Sin(y/b))
}

func moguls(x, y float64) float64 {
	return math.Abs(eggbox(x, y))
}

func saddle(x, y float64) float64 {
	const a = 0.002
	return a*math.Pow(x, 2) - a*math.Pow(y, 2)
}

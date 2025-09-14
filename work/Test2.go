package main

import (
	"fmt"
	"math"
)

var p float64 = 3.14159
var cc float64 = 3.2

func округление(val float64) float64 {
	return math.Round(100*val) / 100
}

func кольцо(d float64) {
	v := (d * p) / cc
	fmt.Println("Количество циклов: ", округление(v), " Уголо поворота стола: ", округление(360/(v*2)))
}

var l, m, c, t float64

func винт(h float64, d float64, v float64) {
	l := p * d
	m := (60 / (l / v)) * 100 / 32
	c := l / v
	t := (c / 1.6) * h
	fmt.Println("l: ", округление(l), " c: ", округление(c), " o: ", округление((m*100)/32), " g: ", округление(t/4), " f: ", округление(t/3))
}

func main() {
	кольцо(120)
	винт(132, 130, 5)
}

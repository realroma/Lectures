package main

import (
	"fmt"
)

var p float64 = 3.14159
var cc float64 = 3.2

func кольцо(d float64) {
	v := (d * p) / cc
	fmt.Println("Количество циклов: ", v, " Уголо поворота стола: ", 360/(v*2))
}

var l, m, c, t float64

func винт(h float64, d float64, v float64) {
	l := p * d
	m := 60 / (l / v)
	c := l / v
	t := (c / 1.6) * h
	fmt.Println("l: ", l, " c: ", c, " o: ", (m*100)/32, " g: ", t/4, " f: ", t/3)
}

func main() {
	кольцо(120)
	винт(132, 130, 5)
}

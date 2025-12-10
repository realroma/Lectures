package main

import (
	"fmt"
	"math"
)

const pi = 3.14

type Circle struct {
	name string
	radius int
}

type Rectangle struct {
	name string
	weight int
	height int
	area int
}

func (struc Circle) Area() {
	area := math.Pow(pi * float64(struc.radius), 2)
	fmt.Println("Area:", area)
}

func (struc Rectangle) Area() {
	fmt.Println("Area:", struc.area)
}

func (struc Circle) Name() {
	fmt.Println("Name:", struc.name)
}

func (struc Rectangle) Name() {
	fmt.Println("Name:", struc.name)
}

func ConstructorCircle (radius int) Circle {
	return Circle {
	name: "Circle",
	radius: radius,
	}
}

func ConstructorRectangle (weight int, height int) Rectangle {
	return Rectangle {
	name: "Rectangle",
	weight: weight,
	height: height,
	area: weight*height*2,
	}
}

func main() {
	circle := ConstructorCircle(5)
	circle.Area()
	circle.Name()
	rectangle := ConstructorRectangle(3, 5)
	rectangle.Name()
	rectangle.Area()
}

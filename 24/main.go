package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: 	x,
		y:	y,
	}
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func sqr(a float64) float64 {
	return a * a
}

func Distance(a, b Point) float64{
	return math.Sqrt(sqr(a.x - b.x) + sqr(a.y - b.y))
}

func main() {
	x1, y1 := 1.2, 2.6
	x2, y2 := -2.0, 1.
	Point1 := New(x1, y1)
	Point2 := New(x2, y2)
	fmt.Println(Distance(*Point1, *Point2))
}
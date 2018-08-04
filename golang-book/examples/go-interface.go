package main

import "math"

type Shape interface {
	area() float64
}
type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(s Shape) float64 {
	return s.area()
}

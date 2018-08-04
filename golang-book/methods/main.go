package main

import (
	"fmt"
	"math"
)

func main() {
	path := Path{
		{1, 2},
		Point{2, 3},
		Point{3, 4},
	}
	fmt.Println(path.Distance())
}

type Point struct{ X, Y float64 }
type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

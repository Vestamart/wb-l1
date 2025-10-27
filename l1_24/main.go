package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y

	dxSquared := dx * dx
	dySquared := dy * dy

	return math.Sqrt(dxSquared + dySquared)
}

func main() {
	p1 := NewPoint(1.0, 2.0)

	p2 := NewPoint(4.0, 6.0)

	distance := p1.Distance(p2)

	fmt.Printf("Координаты P1: (%.1f, %.1f)\n", p1.x, p1.y)
	fmt.Printf("Координаты P2: (%.1f, %.1f)\n", p2.x, p2.y)
	fmt.Printf("Расстояние между P1 и P2: %.2f\n", distance)

	p3 := NewPoint(-1.0, -1.0)
	p4 := NewPoint(2.0, 3.0)
	distance2 := p3.Distance(p4)
	fmt.Printf("\nКоординаты P3: (%.1f, %.1f)\n", p3.x, p3.y)
	fmt.Printf("Координаты P4: (%.1f, %.1f)\n", p4.x, p4.y)
	fmt.Printf("Расстояние между P3 и P4: %.2f\n", distance2)
}

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Нахожу расстояние по теореме пифагора
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Abs(p2.y-p1.y)*math.Abs(p2.y-p1.y) + math.Abs(p1.x-p2.x)*math.Abs(p1.x-p2.x))
}

func main() {
	p1 := NewPoint(0, 3)
	p2 := NewPoint(3, 0)
	fmt.Println(Distance(p1, p2))
}

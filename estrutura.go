package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Vector struct {
	src, dst Point
}

func dist(p1, p2 Point) float64 {
	deltaX := p2.x - p1.x
	deltaY := p2.y - p1.y

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func (v *Vector) norm() float64 {
	return dist(v.src, v.dst)
}

func main() {
	p1 := Point{x: 1, y: 1}
	p2 := Point{x: 0, y: 0}
	fmt.Println(dist(p1, p2))

	p3 := Point{x: 3, y: 4}
	p4 := Point{x: -3, y: -4}
	v := Vector{src: p3, dst: p4}
	fmt.Println(v.norm())
}

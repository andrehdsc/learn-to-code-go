package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("The area of the %T is: %v\n", s, s.area())
}

func main() {
	square := square{
		length: 2,
		width:  2,
	}
	circle := circle{
		radius: 3,
	}
	info(square)
	info(circle)
}

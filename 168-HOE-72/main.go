package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(x float64) func() float64 {
	var k float64
	return func() float64 {
		k++
		return math.Pow(x, k)
	}
}
